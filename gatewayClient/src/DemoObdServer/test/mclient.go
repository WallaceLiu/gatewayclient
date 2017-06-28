package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

const (
	TIMEOUT          = 5
	SERVER_IPANDPORT = "192.168.5.20:8888"
	//SERVER_IPANDPORT = "127.0.0.1:8888"
	CLIENT_COUNT = 15000 //0000
)

type Client struct {
	client   net.Conn
	isAlive  bool
	SendChan chan string
	RecvChan chan string
}

func (c *Client) Connect() bool {

	if c.isAlive {
		return true
	} else {
		var err error
		c.client, err = net.Dial("tcp", SERVER_IPANDPORT)
		if err != nil {
			fmt.Printf("Failure to connect:%s\n", err.Error())
			return false
		}
		c.isAlive = true
	}

	return true
}

func (c *Client) Echo() {

	time.Sleep(2 * time.Second)

	line := <-c.SendChan // 从通道里取出数据，写入，发送给服务端

	fmt.Print("send...")
	log.Printf("%s	request=%s\n", c.client.LocalAddr().String(), line)

	c.client.Write([]byte(line))

	buf := make([]byte, 1024)
	n, err := c.client.Read(buf)
	if err != nil {
		c.RecvChan <- string("server close...")
		log.Printf("server close...")
		c.client.Close()
		c.isAlive = false
		return
	}

	c.RecvChan <- string(buf[0:n])
}

func Work(tc *Client) {
	if !tc.isAlive {
		if tc.Connect() {
			tc.Echo()
		} else {
			<-tc.SendChan
			tc.RecvChan <- string("server close...")
		}
	} else {
		tc.Echo()
	}
}

func create() {
	var tc = &Client{
		SendChan: make(chan string, 1000),
		RecvChan: make(chan string, 1000)}

	if !tc.Connect() {
		return
	}

	cnt := 1
	for {
		go Work(tc)

		tc.SendChan <- fmt.Sprintf("%d", cnt) // 写入通道，准备发送给服务端
		cnt++
		if cnt > 100 {
			cnt = 0
		}
		s := <-tc.RecvChan

		if tc.isAlive {
			fmt.Print("recv...")
			log.Printf("%s	response=%s\n", tc.client.LocalAddr().String(), s)
		} else {
			fmt.Printf("%s\n", s)
		}
	}
}

var logFileName = flag.String("log", "mclient.log", "Log file name")

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
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	fmt.Println("start...")

	for i := 1; i <= CLIENT_COUNT; i++ {
		time.Sleep(time.Millisecond * 10)

		go func() {
			create()
		}()
	}

	<-signals

	time.Sleep(time.Second * 2)
	fmt.Println("...End")
}
