package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	MSG_LENGTH = 1024
)

func handle(c net.Conn) {
	data := make([]byte, MSG_LENGTH)
	defer c.Close()

	for {
		n, err := c.Read(data)
		if err != nil {
			fmt.Print("...")
			log.Printf("read message from failed	%s\n", err.Error())
			return
		}

		fmt.Println(fmt.Sprintf("read %s from %s", string(data[0:n]), c.RemoteAddr().String()))

		_, err = c.Write([]byte(string(data[0:n])))
		if err != nil {
			fmt.Printf("disconnect from server")
			log.Printf("disconnect from server")
			return
		}
	}
}

var logFileName = flag.String("log", "server.log", "Log file name")

func init() {

	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "client start Failed")
		return //os.Exit(-1)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	fmt.Println("start server...")
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Printf("Failure to listen: %s\n", err.Error())
		return
	}

	for {
		if conn, err := l.Accept(); err == nil {
			go handle(conn)
		}
	}
}
