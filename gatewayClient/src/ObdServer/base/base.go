package base

import (
	"fmt"
	"log"
	"net"
	"time"
)

const (
	SERVER_IPANDPORT = "127.0.0.1:8888"
	//SERVER_IPANDPORT = "192.168.5.20:8888"
	CLIENT_NUM = 4000 //64000
	INTERVAL   = 2    // 时间间隔
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

// 写入通道，准备发送给服务端
func WriteChan(c *Client) {

	cnt := 1
	for {
		select {
		case c.SendChan <- fmt.Sprintf("%d", cnt): // 每2秒写一个
			fmt.Print("wc...")
			cnt++
			if cnt > 100 {
				cnt = 0
			}
		}
	}
}

// 从通道里取出数据，写入，发送给服务端
func Sender(c *Client) {

	for {
		time.Sleep(INTERVAL * time.Second)

		select {
		case s := <-c.SendChan:
			_, err := c.client.Write([]byte(s))
			if err != nil {
				fmt.Printf("Write Error.")
				break
			}

			fmt.Print("ws...") // 从chan读取，写入socket
		}
	}
}

func ReadChan(c *Client) {
	for {

		select {
		case s := <-c.RecvChan:
			if c.isAlive {
				fmt.Print("rc...")

			} else {
				fmt.Printf("%s", s)
			}
		}
	}
}

func Revcer(c *Client) {
	buf := make([]byte, 2048)

	for {
		n, err := c.client.Read(buf)
		if err != nil {
			c.RecvChan <- string("server close...")
			log.Printf("server close...")
			c.client.Close()
			c.isAlive = false
			break
		}

		c.RecvChan <- string(buf[0:n])

		fmt.Print("rs...")
	}
}
