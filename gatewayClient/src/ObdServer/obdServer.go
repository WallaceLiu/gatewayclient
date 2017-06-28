package main

import (
	"ObdServer/base"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime"
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

	runtime.Gosched() // 出让时间片

	return cnt
}

func handle(conn net.Conn, userConnMap *UserConnMap) {
	defer conn.Close()

	//reader := bufio.NewReader(conn)

	buf := make([]byte, 2048)

	for {
		rt1 := time.Now()
		rl, errR := conn.Read(buf)
		rt2 := time.Now()
		rt2rt1 := rt2.Sub(rt1).Seconds()
		//byt, errR := reader.ReadByte()

		if errR != nil {
			fmt.Print(fmt.Sprintf("Read Error <%s> FROM <%s>-", errR, conn.RemoteAddr()))
			userConnMap.Del(conn.RemoteAddr().String())
			return
		}

		fmt.Print("r-")

		dt1 := time.Now()
		content := string(buf[0:rl]) + "," + timeConsuming(string(buf[0:rl]))
		dt2 := time.Now()
		dt2dt1 := dt2.Sub(dt1).Seconds()

		wt1 := time.Now()
		_, errW := conn.Write([]byte(content))
		wt2 := time.Now()
		wt2wt1 := wt2.Sub(wt1).Seconds()
		if errW != nil {
			userConnMap.Del(conn.RemoteAddr().String())
			return
		}

		at := wt2.Sub(rt1).Seconds()

		fmt.Print("w-")

		log.Printf("%s	%s	%s	%f	%s	%s	%f	%s	%s	%f	%f\n",
			conn.RemoteAddr().String(),
			rt1.Format(time.RFC3339), rt2.Format(time.RFC3339), rt2rt1-base.INTERVAL,
			dt1.Format(time.RFC3339), dt2.Format(time.RFC3339), dt2dt1,
			wt1.Format(time.RFC3339), wt2.Format(time.RFC3339), wt2wt1,
			at-base.INTERVAL)
	}
}

func cnt(userConnMap *UserConnMap) {

	for {
		time.Sleep(time.Second * 5)
		fmt.Print(userConnMap.Cnt(), "-")
		//log.Printf("%d	-	-	-	-\n", userConnMap.Cnt())
	}
}

func timeConsuming(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

var logFileName = flag.String("log", "obdServer.log", "Log file name")

func init() {

	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "server can not start.")
		return
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
		os.Exit(0)
	}

	var userConnMap = UserConnMap{userConn: make(map[string]net.Conn)}

	address, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf(":%s", os.Args[1]))
	if err != nil {
		fmt.Println("ResolveTCPAddr Error")
		return
	}

	listen, err := net.ListenTCP("tcp", address)

	defer listen.Close()

	if err != nil {
		fmt.Println("Listen Error")
		return
	}

	go cnt(&userConnMap)

	for {
		fmt.Println("Server Listening ...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(fmt.Sprintf("Server Accepting Error，%s", err.Error()))
			continue
		}

		fmt.Printf("Connected from %s\n", conn.RemoteAddr())

		userConnMap.Set(conn.RemoteAddr().String(), conn)

		go handle(conn, &userConnMap)
	}
}
