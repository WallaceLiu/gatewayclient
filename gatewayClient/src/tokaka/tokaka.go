package main

import (
	"config"
	"eq"
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"time"
)

var stopLoop bool = false
var prompt = "Usage: toHBaseFrmKaka %s"
var logFileName = flag.String("log", "producer.log", "Log file name")

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { // Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+C")
	}

	fmt.Println(prompt)
	fmt.Println("Begin...")

	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {

	var (
		totalCount, successCount, failCount = 0, 0, 0
	)

	fmt.Println("Connect Number：", config.MAX_CONNECT_TO_KAKA)
	fmt.Println("KaKa：", strings.Split(config.KAFKA_ADDRS, ","))

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	//cnf := sarama.NewConfig()
	//cnf.Producer.Flush.Messages = 10
	//cnf.Producer.Return.Successes = true
	//cnf.Producer.Retry.Max = 3
	//cnf.Producer.Retry.Backoff = 0
	//cnf.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	//cnf.Producer.RequiredAcks = sarama.WaitForAll
	//cnf.Producer.Partitioner = sarama.NewRandomPartitioner

	for i := config.MIN_TO_KAKA; i < config.MAX_CONNECT_TO_KAKA; i++ {
		totalCount++

		producer, err := sarama.NewAsyncProducer(strings.Split(config.KAFKA_ADDRS, ","), nil)

		if nil != err {
			failCount++
			fmt.Println("Connect SEQ=", i, " Connect Kaka Fail, ERR=", err.Error())
		}

		successCount++

		time.Sleep(time.Millisecond * 500)

		eq := eq.NewEq(int64(i))

		go myProducer(producer, eq, int64(i))
	}

	<-signals

	stopLoop = true

	fmt.Println("----------Total=", totalCount, ",Success=", successCount, ",Fail=", failCount)

	time.Sleep(time.Second * 5)

	fmt.Println("End.")

}

func myProducer(p sarama.AsyncProducer, eq *eq.Eq, seq int64) {

	var num = 1

	//defer closeProducer(p)

	defer func() {
		if p != nil {
			p.AsyncClose()
		}
	}()

	var (
		totalMsg = 0
	)

	msg := &sarama.ProducerMessage{}
	msg.Topic = config.KAFKA_TOPIC
	msg.Partition = int32(-1)
	msg.Key = sarama.StringEncoder("key")
	msg.Value = sarama.ByteEncoder("hello world!")

	for !stopLoop {
		eq = eq.SetEq()

		msg.Key = sarama.StringEncoder(eq.Vin)
		msg.Value = sarama.ByteEncoder(eq.StrMsg)
		//msg.Value = sarama.ByteEncoder(eq.JsonMsg)

		fmt.Printf(">>> Produce SEQ=%d,msg=%s\n", seq, msg.Value)
		log.Printf(">>> Produce SEQ=%d,msg=%s\n", seq, msg.Value)

		num = num + 1

		p.Input() <- msg

		totalMsg++

		time.Sleep(time.Millisecond * 10)
	}
}

func closeProducer(p sarama.AsyncProducer) {
	//p.AsyncClose()
	if p != nil {
		var wg sync.WaitGroup

		p.AsyncClose()

		wg.Add(2)
		go func() {
			for _ = range p.Successes() {
				fmt.Println("Unexpected message on Successes()")
			}
			wg.Done()
		}()
		go func() {
			for msg := range p.Errors() {
				fmt.Println(msg.Err)
			}
			wg.Done()
		}()
		wg.Wait()
	}
}
