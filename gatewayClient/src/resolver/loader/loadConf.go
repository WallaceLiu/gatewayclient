package loader

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"resolver/base"
	"runtime"
	"strings"
)

var workers = runtime.NumCPU()
var ProtocolMapBytes = make(map[string]protocolBytes)

type protocolBytes struct {
	filename string         // 协议文件名
	bytes    []byte         // 协议字节数组
	p        *base.Protocol // 协议结构体
}

type job struct {
	filename string
	results  chan<- protocolBytes
}

func (job job) Do() {
	if content, err := ioutil.ReadFile("../conf/" + job.filename); err != nil {
		log.Fatal(err)
	} else {
		var p *base.Protocol
		if err := xml.Unmarshal(content, &p); err != nil {
			log.Fatal(err)
		}

		job.results <- protocolBytes{job.filename, content, p}
	}
}

func LoadConf() {
	filenames := GetProtocolFiles()
	jobs := make(chan job, workers)
	results := make(chan protocolBytes, minimum(100, len(filenames)))
	done := make(chan struct{}, workers)

	go addJobs(jobs, filenames, results)

	for i := 0; i < workers; i++ {
		go doJobs(done, jobs)
	}

	waitAndProcessResults(done, results)
}
func GetProtocolFiles() []string {
	files := make([]string, 0, 2)

	if dir, err := ioutil.ReadDir("../conf/"); err != nil {
		log.Fatal(err)
		return nil
	} else {
		for _, file := range dir {
			if file.IsDir() {
				continue
			}
			files = append(files, file.Name())
		}
		return files
	}
}

func addJobs(jobs chan<- job, filenames []string, results chan<- protocolBytes) {
	for _, filename := range filenames {
		jobs <- job{filename, results} // 添加不进去时，会阻塞
	}
	close(jobs) // 所有文件添加到jobs后，并且被取出去，会关闭jobs通道
}

func doJobs(done chan<- struct{}, jobs <-chan job) {
	for job := range jobs {
		job.Do()
	}

	done <- struct{}{}
}

func waitAndProcessResults(done <-chan struct{}, results <-chan protocolBytes) {

	for working := workers; working > 0; {
		select { // Blocking
		case result := <-results:
			ProtocolMapBytes[strings.TrimSuffix(result.filename, ".xml")] = result
		case <-done:
			working--
		}
	}
DONE:
	for {
		select { // Nonblocking
		case result := <-results:
			ProtocolMapBytes[strings.TrimSuffix(result.filename, ".xml")] = result
		default:
			break DONE
		}
	}

}

func minimum(x int, ys ...int) int {
	for _, y := range ys {
		if y < x {
			x = y
		}
	}
	return x
}
