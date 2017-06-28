package common

import (
	"log"
	"strconv"
	"testing"
)

func TestBytesSliceToInt8(t *testing.T) {

}

func TestBytesSliceToInt16(t *testing.T) {
	var b = []byte{22, 8}
	log.Println(BytesSliceToInt16(b))

}

func TestBytesSliceToUInt8(t *testing.T) {
	var b2 = []byte{255}
	log.Println(BytesSliceToUInt8(b2))
}

func TestBytesSliceToInt32(t *testing.T) {
	i := "5640"
	j, err := strconv.Atoi(i)
	if err != nil {
		log.Println("err")
	} else {
		log.Printf("%x", j)
	}

}

func TestBytesSliceToInt64(t *testing.T) {

}
