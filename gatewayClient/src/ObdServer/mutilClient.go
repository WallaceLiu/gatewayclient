package main

import (
	"ObdServer/base"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func handler() {
	var tc = &base.Client{
		SendChan: make(chan string, 1000),
		RecvChan: make(chan string, 1000)}

	if !tc.Connect() {

		return
	}

	go base.WriteChan(tc)
	go base.Sender(tc)

	go base.ReadChan(tc)
	go base.Revcer(tc)
}

var logFileName = flag.String("log", "mutilClient.log", "Log file name")

//var conf *cnf.ObdServerConfig

func init() {

	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "client start Failed")
		return
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//conf = cnf.GetConfig()
}

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for i := 1; i <= base.CLIENT_NUM; i++ {
		time.Sleep(time.Millisecond * 10)

		go func() {
			handler()
		}()
	}

	<-signals

	time.Sleep(time.Second * 2)
	fmt.Println("...End")
}
