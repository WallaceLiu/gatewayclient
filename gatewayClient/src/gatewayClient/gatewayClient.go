package main

import (
	"config"
	"eq"
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"
	"time"
)

var stopLoop bool = false
var prompt = "Usage: gatewayClient <ip:port> %s"

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { // Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+C")
	}

	fmt.Println(prompt)
	fmt.Println("Begin...")
}

func main() {

	svrAddr := os.Args[1]
	fmt.Println("Server Address：", svrAddr)

	var (
		totalCount, successCount, failCount = 0, 0, 0
	)

	fmt.Println("TCP/IP connect Number：", config.MAX_CONNECT)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for i := config.MIN; i < config.MAX; i++ {
		totalCount++

		conn, err := net.DialTimeout("tcp", svrAddr, time.Second)

		if nil != err {
			failCount++
			fmt.Println("Connect SEQ=", i, " Connect Server Fail, ERR=", err.Error())
		}

		successCount++
		// 每100哈毫秒建立一个连接
		//time.Sleep(time.Millisecond * 100)
		time.Sleep(time.Second * 2)

		eq := eq.NewEq(int64(i))

		go client(conn, eq, int64(i))

		go recMessage(conn, int64(i))

		fmt.Println("Connect SEQ=", i, " Connect Server Success...")
	}

	<-signals

	stopLoop = true

	fmt.Println("----------Total=", totalCount, ",Success=", successCount, ",Fail=", failCount)

	time.Sleep(time.Second * 5)

	fmt.Println("End.")

}

func client(conn net.Conn, eq *eq.Eq, seq int64) {
	var (
		totalMsg, sendSucccess, sendFail = 0, 0, 0
	)

	defer conn.Close()

	for !stopLoop {
		eq = eq.SetEq()

		msgByte := eq.BytesMsg
		fmt.Println(fmt.Sprintf(">>> Send SEQ=%d BYTE=%x", seq, msgByte)) //.Println(">>> Send SEQ=", seq, " BYTE=", msgByte, "...")

		_, err := conn.Write(msgByte)

		totalMsg++

		if err != nil {
			sendFail++
			fmt.Println(">>> SEQ=", seq, " Send Err：", err.Error())
		} else {
			sendSucccess++
		}

		//time.Sleep(time.Millisecond * 100)
		time.Sleep(time.Second * 1)
	}

	fmt.Println("===> SEQ=", seq,
		",Total=", totalMsg,
		",Success=", sendSucccess,
		",Fail=", sendFail)
}

func recMessage(conn net.Conn, seq int64) {
	recSuccess, recFail := 0, 0

	defer conn.Close()

	for !stopLoop {
		recBuf := make([]byte, 100)

		_, err := conn.Read(recBuf)
		if err != nil {
			recFail++
			fmt.Println("<<< Rec SEQ=", seq, " Err：", err.Error())
		} else {
			recSuccess++
		}

		fmt.Println("<<< Rec SEQ=", seq, " Message=", recBuf[0:2], "...")
	}

	fmt.Println("<=== Rec SEQ=", seq,
		",Rec Success=", recSuccess,
		",Rec Fail=", recFail)
}
