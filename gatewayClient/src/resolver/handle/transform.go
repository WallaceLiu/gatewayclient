package handle

import (
	"common"
	"fmt"
	"log"
	"resolver/base"
	"resolver/util"
)

// 字节数组转换
func transform(bArr []byte, start *int, segment *base.Segment) (value string) {
	var t []byte
	var s, l = *start, 0

	switch util.ConvertTranformStrTo(segment.Transform) {

	case base.ASCII:
		t = bArr[*start : *start+segment.Len]
		value = common.ByteSliceToStr(t)
		l = segment.Len
		*start += segment.Len

	// 字符串
	// 例如：
	// 字节数组 [53 49 50 48 52 48 50 55 51]
	// 为
	// 字符串 512040273
	case base.STRING:
		i := *start
		for {
			if bArr[i] == base.ASCII_END {
				break
			}
			i++
		}
		t = bArr[*start:i]
		value = common.ByteSliceToStr(t)
		l = i - *start + 1
		*start = i + 1

	case base.HEX_STRING:
		value = ""
		*start += 0
	case base.COMMA:
		i, j := 1, *start
		for {
			if bArr[j] == base.ASCII_COMMA {
				break
			}
			i++
		}
		t = bArr[*start : *start+i-1]
		value = common.ByteSliceToStr(t)
		l = i
		*start = i + 1

	case base.HEX:
		switch segment.Len {
		case 1:
			t = bArr[*start : *start+segment.Len]
			value = fmt.Sprintf("%x", t)
		default:
			value = "FF"
		}
		l = segment.Len
		*start += segment.Len

	// 整型

	case base.NUMERAL:
		t = bArr[*start : *start+segment.Len]
		value = fmt.Sprintf("%d", common.BytesSliceToUInt8(t))
		l = segment.Len
		*start += segment.Len
	// 一字节转换无符号整型
	// 例如
	// 字节数组 [255]
	// 为
	// 无符号整型 255
	case base.NUMERAL_TO_UINT8:
		t = bArr[*start : *start+segment.Len]
		value = fmt.Sprintf("%d", common.BytesSliceToUInt8(t))
		l = segment.Len
		*start += segment.Len

	case base.NUMERAL_TO_INT8:
		t = bArr[*start : *start+segment.Len]
		value = fmt.Sprintf("%d", common.BytesSliceToInt8(t))
		l = segment.Len
		*start += segment.Len

	case base.NUMERAL_TO_UINT16:
		t = bArr[*start : *start+segment.Len]
		value = fmt.Sprintf("%d", common.BytesSliceToUInt16(t))
		l = segment.Len
		*start += segment.Len
	// 两字节转换整型
	// 例如
	// 字节数组 [22, 8]
	// 为
	// 两个字节整型 5640
	case base.NUMERAL_TO_INT16:
		t = bArr[*start : *start+segment.Len]
		value = fmt.Sprintf("%d", common.BytesSliceToInt16(t))
		l = segment.Len
		*start += segment.Len

	case base.NUMERAL_TO_UINT32:
		t = bArr[*start : *start+segment.Len]
		value = fmt.Sprintf("%d", common.BytesSliceToUInt32(t))
		l = segment.Len
		*start += segment.Len
	case base.NUMERAL_TO_INT32:
		t = bArr[*start : *start+segment.Len]
		value = fmt.Sprintf("%d", common.BytesSliceToInt32(t))
		l = segment.Len
		*start += segment.Len

	case base.NUMERAL_TO_UINT64:
		t = bArr[*start : *start+segment.Len]
		value = fmt.Sprintf("%d", common.BytesSliceToUInt64(t))
		l = segment.Len
		*start += segment.Len
	case base.NUMERAL_TO_INT64:
		t = bArr[*start : *start+segment.Len]
		value = fmt.Sprintf("%d", common.BytesSliceToInt64(t))
		l = segment.Len
		*start += segment.Len

	default:
		value = ""
		l = 0
		*start += 0
	}

	log_head := fmt.Sprint("(", segment.Id, ",", segment.Title, ")")
	log_content := fmt.Sprint(",(", t, ",", len(bArr), ")", ",type=", segment.Ty, ",trans=", segment.Transform, ",value=", value)
	var log_before string
	if util.ConvertTranformStrTo(segment.Transform) == base.STRING {
		log_before = fmt.Sprint(",(", s, ",-)")
	} else {
		log_before = fmt.Sprint(",(", s, ",", l, ")")
	}
	log_after := fmt.Sprint("(", *start, ",", l, ")")

	log.Println("Parse Segment...", log_before, "->", log_after, log_head, log_content)

	return value
}
