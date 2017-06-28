package common

import (
	"log"
	"testing"
)

func Test_Int32ToBytesSlice(t *testing.T) {
	log.Println(Int32ToBytesSlice(100))
}

func Test_Int16ToBytesSlice(t *testing.T) {
	log.Println(Int16ToBytesSlice(200))
}

func Test_Float64ToBytesSlice(t *testing.T) {
	log.Println(Float64ToBytesSlice(100))
}

func Test_UInt16ToBytesSlice(t *testing.T) {
	log.Println(UInt16ToBytesSlice(5638))
}

func Test_ByteSliceNot(t *testing.T) {
	srcB := []byte{0x55, 0xaa}
	destB := ByteSliceNot(srcB)
	log.Println("%x", srcB[0], "%x", destB[0])
	log.Println("%x", srcB[1], "%x", destB[1])
}

func Test_StrToByteSliceNoB0(t *testing.T) {
	log.Println(StrToByteSliceNoB0(",")[0])
}
