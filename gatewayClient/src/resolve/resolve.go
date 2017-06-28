package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"resolver/test"
)

var logFileName = flag.String("log", "data.log", "Log file name")

func init() {

	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "client start Failed")
		return
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	//signals := make(chan os.Signal, 1)
	//signal.Notify(signals, os.Interrupt)
	var stop = make(chan bool)

	test.SplitMsg([]byte{170, 85}, stop)

	<-stop

	fmt.Println("End.")
}
