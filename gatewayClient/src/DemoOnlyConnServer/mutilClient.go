package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	//"os/signal"
	//"runtime"
	"os/signal"
	"sync"
	"time"
)

const (
	SERVER_IPANDPORT = "192.168.5.20:8888" // //"127.0.0.1:8888"
	CLIENT_NUM       = 64000               //64000
)

func create(address string) {
	addr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		log.Fatalf("ResolveTCPAddr	-	%s\n", err.Error())
		return
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("DialTCP Err：", err.Error())
		log.Fatalf("DialTCP	-	%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("%s	Connected	%s", conn.LocalAddr().String(), conn.RemoteAddr().String()))

	//Add(lock)

	buf := make([]byte, 1024)

	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(fmt.Sprintf("%s Lost,", conn.RemoteAddr().String()), err.Error())
			conn.Close()
			log.Fatalf("%s Read	-	-	-	%s\n", conn.LocalAddr().String(), err.Error())
			//Sub(lock)
			break
		}
	}
}

var logFileName = flag.String("log", "mutilClient.log", "Log file name")

func init() {

	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "client start Failed")
		return //os.Exit(-1)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

var counter int = 0
var isMaxCnt = false

func Add(lock *sync.Mutex) {
	lock.Lock() // 给变量counter加锁
	counter++
	if counter >= CLIENT_NUM {
		isMaxCnt = true
	}
	fmt.Println("counter =", counter)
	lock.Unlock() // 解锁
}
func Sub(lock *sync.Mutex) {
	lock.Lock() // 给变量counter加锁
	if counter > 0 {
		counter--
	}
	fmt.Println("counter =", counter)
	lock.Unlock() // 解锁
}

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	fmt.Println("Only Connect...")
	//lock := &sync.Mutex{}

	for i := 1; i <= CLIENT_NUM; i++ {
		time.Sleep(time.Millisecond * 100)

		go func() {
			create(SERVER_IPANDPORT)
		}()

	}

	//for {
	//	lock.Lock()
	//
	//	c := counter
	//
	//	lock.Unlock()
	//
	//	runtime.Gosched() // 出让时间片
	//
	//	// 10个协程，记数到10退出程序
	//	if isMaxCnt && c <= 0 {
	//		break
	//	}
	//}

	<-signals

	time.Sleep(time.Second * 2)
	fmt.Print("...End")
}
