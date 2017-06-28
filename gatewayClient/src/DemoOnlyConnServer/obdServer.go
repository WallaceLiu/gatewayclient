package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type UserConnMap struct {
	userConn map[string]net.Conn
	sync.RWMutex
}

func (s *UserConnMap) Set(key string, conn net.Conn) {
	s.Lock()
	s.userConn[key] = conn
	s.Unlock()
}

func (s *UserConnMap) Del(key string) {
	s.Lock()
	delete(s.userConn, key)
	s.Unlock()
}

func (s UserConnMap) Cnt() (cnt int) {
	s.Lock()
	cnt = len(s.userConn)
	s.Unlock()
	return cnt
}

func handle(conn net.Conn, userConnMap *UserConnMap) {
	defer func() {
		fmt.Println(fmt.Sprintf("%s Closed", conn.RemoteAddr().String()))
		log.Println(fmt.Sprintf("%s Closed", conn.RemoteAddr().String()))
		conn.Close()
	}()

	reader := bufio.NewReader(conn)

	for {
		_, err := reader.ReadByte()

		if err != nil {
			fmt.Println(fmt.Sprintf("Reading Error <%s> FROM <%s>", err, conn.RemoteAddr()))
			userConnMap.Del(conn.RemoteAddr().String())
			return
		}
	}
}

func cnt(userConnMap *UserConnMap) {

	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Number of Connections is：", userConnMap.Cnt())
		log.Println(fmt.Sprintf("%d	Connections", userConnMap.Cnt()))
	}
}

var logFileName = flag.String("log", "obdServer.log", "Log file name")

func init() {

	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "obdServer start Failed")
		return //os.Exit(-1)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <Server Port>\n", filepath.Base(os.Args[0]))
		fmt.Printf("usage: %s <Server IP Address>:<Server Port>\n", filepath.Base(os.Args[0]))
		fmt.Printf("    eg: chatServer 8888\n")
		fmt.Printf("    eg: chatServer 192.168.0.74:8888\n")
		return //os.Exit(0)
	}

	var userConnMap = UserConnMap{userConn: make(map[string]net.Conn)}

	address, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf(":%s", os.Args[1]))
	if err != nil {
		fmt.Println(fmt.Sprintf("ResolveTCPAddr,%s", err.Error()))
		return //os.Exit(-1)
	}

	listen, err := net.ListenTCP("tcp", address)

	defer listen.Close()

	if err != nil {
		fmt.Println(fmt.Sprintf("ListenTCP，%s", err.Error()))
		os.Exit(-1)
	}

	go cnt(&userConnMap)

	for {
		fmt.Println("Listening ...")

		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(fmt.Sprintf("Accepting Error，%s", err.Error()))
			log.Println(fmt.Sprintf("Accepting Error，%s", err.Error()))
			continue
		}

		clientName := conn.RemoteAddr().String()

		fmt.Println(fmt.Sprintf("%s Connected", clientName))
		fmt.Println("Acceped ...")

		log.Println(fmt.Sprintf("%s Connected", clientName))

		userConnMap.Set(conn.RemoteAddr().String(), conn)

		go handle(conn, &userConnMap)
	}
}
