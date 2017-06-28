package test

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

/*
	写入通道				切分消息		解析消息
        c1 READ_FILE_BUF		c2
 	goroutine      			goroutine	goroutine
 	|------------------|		|----|		|    |
 	|READ_FILE_BUF_SIZE|		|msg |		|    |
 	|READ_FILE_BUF_SIZE|		|msg |		|    |
读取文件 	|READ_FILE_BUF_SIZE| s<-c1	|msg | msg<-c2	|    |
 	|READ_FILE_BUF_SIZE|		|msg |		|    |
 	|..................|		|msg |		|    |
 	|------------------|		|----|		|----|

*/

const (
	TEST_FILE_PATH = "../../resolver/testfile/zbinary_1m.dat" // 测试文件路径
	//READ_FILE_BUF_SIZE  = 1024                                  // 每次读取文件缓存区大小
	READ_FILE_BUF_SIZE  = 1024 // 每次读取文件缓存区大小
	MSG_CHAN_BUF_SIZE   = 1000 // 消息通道缓存大小
	MSG_SPLIT_CHAN_SIZE = 1000 // 分隔后消息通道大小
)

type message struct {
	msg    []byte
	length int
}

var readTotalByte = 0

func SplitMsg(seq []byte, stop chan bool) {
	file, err := os.Open(TEST_FILE_PATH)

	if err != nil {
		log.Println(err)
		return
	}

	var cBArr = make(chan []byte, MSG_CHAN_BUF_SIZE)
	var cMsg = make(chan message, MSG_SPLIT_CHAN_SIZE)

	go cat(file, cBArr, stop)

	go split(seq, cBArr, cMsg)

	var i = 1
	for m := range cMsg {
		//time.Sleep(1 * time.Second)
		//fmt.Println(i, ",(", len(byt), ",", len(cMsg), "),", c, ",read total byte from file=", totalByte)
		fmt.Println(i, ",(", len(cBArr), ",", len(cMsg), "),", m)
		//log.Println(i, ",(", len(cBArr), ",", len(cMsg), "),", m)
		i++
	}

	//go wrapper(cMsgByte, cMsg)

	//go handle(cMsg)
}

/* 读取测试文件
 */
func cat(f *os.File, byt chan []byte, stop chan bool) {

	buf := make([]byte, READ_FILE_BUF_SIZE)
	for {

		switch n, err := f.Read(buf[:]); true {
		case n < 0:
			log.Println(os.Stderr, "Read Error：%s\n", err.Error())
			break
		case n == 0: // EOF
			break
		case n > 0:
			readTotalByte += n
			log.Println("Read From File：", buf[0:n])
			byt <- buf[0:n]
		}
	}

	stop <- true
}

/* 分隔消息
 */
func split(seq []byte, cBArr chan []byte, cMsg chan message) {

	var s []byte
	for {
		b := <-cBArr
		log.Println("Get：", b)
		s = append(s, b...)
		log.Println("Append：", s)

		r, idx := splitMsgByte(s, seq)
		log.Println("Split：", r)
		if idx > 0 {
			d := new(message)
			d.msg = r
			d.length = len(r)
			//msgTotalByte += d.length
			cMsg <- *d
			s = s[idx:]
		}
	}
}

/* 添加额外封装 data 的操作
 */
func wrapper(cMsgByte chan []byte, cMsg chan message) {
	for {
		//select {
		//case s := <-cMsgByte:
		//	d := new(data)
		//	d.msg = s
		//	d.length = len(s)
		//	cMsg <- *d
		//}
		s := <-cMsgByte
		d := new(message)
		d.msg = s
		d.length = len(s)
		cMsg <- *d
	}
}

/* 解析消息
 */
func handle(cMsg chan message) {
	for {
		//select {
		//case m, err := <-cMsg:
		//	if err ==false{
		//		fmt.Println("handle ", err)
		//	} else {
		//		fmt.Println("handle：", m.length)
		//		log.Println(time.Now(), ",", m.length, " byte")
		//	}
		//}

		m := <-cMsg
		//fmt.Println(m)
		if m.length == 0 {
			continue
		}
		fmt.Println("handle：", m.length)
		log.Println(time.Now(), ",", m.length, " byte")

		//if len(cMsg) >= 2 {
		//	stop <- true
		//	return
		//}
	}
}

/* 分隔消息
 */
func splitMsgByte(s []byte, sep []byte) (subBy []byte, idx int) {
	return splitMsgBytes(s, sep, 0, 0)
}

/* 分隔消息
 */
func splitMsgBytes(s []byte, sep []byte, start int, length int) (subBy []byte, idx int) {
	if len(s[start:]) >= len(sep) {
		idx = bytes.Index(s[start+len(sep):], sep)
		subBy = s[start : idx+len(sep)+length]

		if idx > 0 {
			return subBy, idx + len(sep) + length
		}
		return s[start:], 0
	}

	return s[start:], 0
}
