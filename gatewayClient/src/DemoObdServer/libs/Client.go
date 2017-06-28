package tcplotus

import (
	"log"
	"net"
)

type Client struct {
	client   net.Conn
	isAlive  bool
	SendChan chan *Request
	RecvStr  chan string
}

type Request struct {
	reqId      int
	reqContent string
	rspChan    chan<- string // write only for client
}

func (c *Client) Connect() bool {
	if c.isAlive {
		return true
	} else {
		var err error
		c.client, err = net.Dial("tcp", proxy_server)
		if err != nil {
			return false
		}
		c.isAlive = true
		log.Println("connect to " + proxy_server)
	}
	return true
}
