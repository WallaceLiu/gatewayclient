package test

import (
	"common"
	"log"
	"os"
	"testing"
)

// 额外测试，不要执行，耗时
func TestSplitMsgByte(t *testing.T) {
	file, err := os.Open("testfile/zbinary_1m.dat")
	if err != nil {
		log.Println(err)
	}
	by := common.Cat(file)
	log.Println(len(by))

	//bb1, idx1 := splitMsgByte(b, []byte{170, 85}, 0, 0)
	//fmt.Println(bb1)
	//fmt.Println(idx1)

	//bb2, idx2 := splitMsgByte(b, []byte{170, 85}, idx1, len(bb1))
	//fmt.Println(bb2)
	//fmt.Println(idx2)
	//
	//bb3, idx3 := splitMsgByte(b, []byte{170, 85}, idx2, len(bb2)+len(bb1))
	//fmt.Println(bb3)
	//fmt.Println(idx3)
	//
	//bb4, idx4 := splitMsgByte(b, []byte{170, 85}, idx3, len(bb3)+len(bb2)+len(bb1))
	//fmt.Println(bb4)
	//fmt.Println(idx4)
	//
	//bb5, idx5 := splitMsgByte(b, []byte{170, 85}, idx4, len(bb4)+len(bb3)+len(bb2)+len(bb1))
	//fmt.Println(bb5)
	//fmt.Println(idx5)
}
