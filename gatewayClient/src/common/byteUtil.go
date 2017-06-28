package common

import (
	"bytes"
	"encoding/binary"
)

// 字符串转换字节切片
func StrToByteSlice(str string) []byte {
	return append(bytes.NewBufferString(str).Bytes(), []byte{0}...)
}

//
func StrToByteSliceNoB0(str string) []byte {
	return bytes.NewBufferString(str).Bytes()
}

// float 64 转换成字节切片
func Float64ToBytesSlice(f float64) []byte {
	t := float64(f)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, t)
	return bytesBuffer.Bytes()
}

//
func Int32ToBytesSlice(n int32) []byte {
	t := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, t)
	return bytesBuffer.Bytes()
}

//
func Int16ToBytesSlice(n int16) []byte {
	t := int16(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, t)
	return bytesBuffer.Bytes()
}

//
func UInt16ToBytesSlice(n uint16) []byte {
	t := uint16(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, t)
	return bytesBuffer.Bytes()
}

//
func ByteSliceNot(bytes []byte) []byte {
	res := make([]byte, len(bytes))

	for i, b := range bytes {
		res[i] = ^b
	}

	return res
}

func ByteRemove(b []byte, idx int) []byte {
	slice := []byte{}
	if idx == 0 {
		slice = append(slice, b[1:]...)
	} else if idx == len(b)-1 {
		slice = append(slice, b[0:len(b)-1]...)
	} else {
		slice = append(slice, b[0:idx-1]...)
		slice = append(slice, b[idx:]...)
	}
	return slice
}
