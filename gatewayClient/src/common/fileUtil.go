package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(filePath string) []byte {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err.Error())
		return []byte{}
	}
	return b
}

func Cat(f *os.File) []byte {
	var payload []byte
	for {
		buf := make([]byte, 1024)
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "reading error：%s\n", err.Error())
			return []byte{}
		case nr == 0: // EOF
			return payload
		case nr > 0:
			payload = append(payload, buf...)
		}
	}
}
