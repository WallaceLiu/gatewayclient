package tcplotus

import (
	"encoding/json"
	"log"
	"net"
	"strconv"
	"time"
)

const (
	proxy_timeout = 5
	proxy_server  = "127.0.0.1:8888"
	msg_length    = 1024
)

//store request map
var requestMap map[int]*Request

func Send(c *Client) {

	//store reqId and reqContent
	senddata := make(map[string]string)
	for {
		if !c.isAlive {
			time.Sleep(1 * time.Second)
			c.Connect()
		}

		if c.isAlive {
			req := <-c.SendChan

			//construct request json string
			senddata["reqId"] = strconv.Itoa(req.reqId)
			senddata["reqContent"] = req.reqContent
			sendjson, err := json.Marshal(senddata)
			if err != nil {
				continue
			}

			_, err = c.client.Write([]byte(sendjson))
			if err != nil {
				c.RecvStr <- string("server close...")
				c.client.Close()
				c.isAlive = false
				log.Println("disconnect from " + proxy_server)
				continue
			}
			//log.Println("Write to proxy server: " + string(sendjson))
		}
	}
}

func Recv(c *Client) {
	buf := make([]byte, msg_length)
	recvdata := make(map[string]string, 2)
	for {
		if !c.isAlive {
			time.Sleep(1 * time.Second)
			c.Connect()
		}

		if c.isAlive {
			n, err := c.client.Read(buf)
			if err != nil {
				c.client.Close()
				c.isAlive = false
				log.Println("disconnect from " + proxy_server)
				continue
			}
			//log.Println("Read from proxy server: " + string(buf[0:n]))

			if err := json.Unmarshal(buf[0:n], &recvdata); err == nil {
				reqidstr := recvdata["reqId"]
				if reqid, err := strconv.Atoi(reqidstr); err == nil {
					req, ok := requestMap[reqid]
					if !ok {
						continue
					}
					req.rspChan <- recvdata["resContent"]
				}
				continue
			}
		}
	}
}

//one handle per request
func handle(conn *net.TCPConn, id int, tc *Client) {

	data := make([]byte, msg_length)
	handleProxy := make(chan string)
	request := &Request{reqId: id, rspChan: handleProxy}
	requestMap[id] = request // 添加请求

	for {
		n, err := conn.Read(data)
		if err != nil {
			log.Println("disconnect from " + conn.RemoteAddr().String())
			conn.Close()
			delete(requestMap, id)
			return
		}

		//fmt.Println(n)

		request.reqContent = string(data[0:n])

		//send to server
		select {
		case tc.SendChan <- request:
		case <-time.After(proxy_timeout * time.Second):
			//proxyChan <- &Request{cancel: true, reqId: id}
			_, err = conn.Write([]byte("server send timeout."))
			if err != nil {
				conn.Close()
				delete(requestMap, id)
				return
			}
			continue
		}

		//read from server
		select {
		case rspContent := <-handleProxy:
			_, err := conn.Write([]byte(rspContent))
			if err != nil {
				conn.Close()
				delete(requestMap, id)
				return
			}
		case <-time.After(proxy_timeout * time.Second):
			_, err = conn.Write([]byte("server recv timeout."))
			if err != nil {
				conn.Close()
				delete(requestMap, id)
				return
			}
			continue
		}
	}
}

func TcpLotusMain(ip string, port int) {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP(ip), Port: port, Zone: ""})
	if err != nil {
		log.Fatalln("listen port error")
		return
	}

	log.Println("start server " + ip + " " + strconv.Itoa(port))

	defer listen.Close()

	var c Client
	c.SendChan = make(chan *Request, 1000) // 创建1000个客户端发送通道
	c.RecvStr = make(chan string)          // 创建1个客户端接收通道
	c.Connect()                            // 尝试连接服务端

	go Send(&c)
	go Recv(&c)

	requestMap = make(map[int]*Request) // 客户端请求map
	var id int = 0
	for {

		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Println("receive connection failed")
			continue
		}

		id++
		log.Println("connected from " + conn.RemoteAddr().String())

		go handle(conn, id, &c)
	}
}
