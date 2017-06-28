package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

// 客户端发送线程
func chat(conn net.Conn) {

	var input string
	username := conn.LocalAddr().String()
	for {

		fmt.Scanln(&input)
		if input == "/quit" {
			fmt.Println("ByeBye..")
			conn.Close()
			os.Exit(0)
		}

		lens, err := conn.Write([]byte(username + " Say :::" + input))
		fmt.Println(lens)
		if err != nil {
			fmt.Println(err.Error())
			conn.Close()
			break
		}
	}
}

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <Server IP Address>:<Server Port>\n", filepath.Base(os.Args[0]))
		fmt.Printf("    eg: chatClient 127.0.0.1:8888\n")
		os.Exit(0)
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp4", os.Args[1])
	if err != nil {
		fmt.Println(fmt.Sprintf("ResolveTCPAddr,%s", err.Error()))
		os.Exit(-1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(fmt.Sprintf("DialTCP,%s", err.Error()))
		os.Exit(-1)
	}

	go chat(conn)

	buf := make([]byte, 1024) //开始客户端轮训

	for {

		lenght, err := conn.Read(buf)
		if err != nil {
			fmt.Println(fmt.Sprintf("Connection,%s", err.Error()))
			os.Exit(-1)
		}

		fmt.Println(string(buf[0:lenght]))
	}
}
