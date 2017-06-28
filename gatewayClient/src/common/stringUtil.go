package common

import "bytes"

// 字节数组切片转换成字符串
func ByteSliceToStr(bArr []byte) string {
	return bytes.NewBuffer(bArr).String()
}
