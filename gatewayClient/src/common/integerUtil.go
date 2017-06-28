package common

import (
	"bytes"
	"encoding/binary"
)

// slice -> uint8 int8
func BytesSliceToInt8(bSlice []byte) int8 {
	bytesBuffer := bytes.NewBuffer(bSlice)
	var tmp int8
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int8(tmp)
}

func BytesSliceToUInt8(bSlice []byte) uint8 {
	bytesBuffer := bytes.NewBuffer(bSlice)
	var tmp uint8
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return uint8(tmp)
}

// slice -> uint16 int16
func BytesSliceToUInt16(bSlice []byte) uint16 {
	bytesBuffer := bytes.NewBuffer(bSlice)
	var tmp uint16
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return uint16(tmp)
}

func BytesSliceToInt16(bSlice []byte) int16 {
	bytesBuffer := bytes.NewBuffer(bSlice)
	var tmp int16
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int16(tmp)
}

// slice -> uint32 int32
func BytesSliceToUInt32(bSlice []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(bSlice)
	var tmp uint32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return uint32(tmp)
}

func BytesSliceToInt32(bSlice []byte) int32 {
	bytesBuffer := bytes.NewBuffer(bSlice)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int32(tmp)
}

// slice -> uint64 int64
func BytesSliceToUInt64(bSlice []byte) uint64 {
	bytesBuffer := bytes.NewBuffer(bSlice)
	var tmp uint64
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return uint64(tmp)
}

func BytesSliceToInt64(bSlice []byte) int64 {
	bytesBuffer := bytes.NewBuffer(bSlice)
	var tmp int64
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int64(tmp)
}

