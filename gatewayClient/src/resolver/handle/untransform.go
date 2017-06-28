package handle

import (
	"resolver/base"
	"resolver/util"
)

// 字节数组转换
func unTransform(segment *base.Segment) (value []byte) {
	switch util.ConvertTranformStrTo(segment.Transform) {
	case base.ASCII:
		value = []byte{0} //common.ByteSliceToStr(bArr[start:length])
		// 字符串
	case base.COMMA:
		value = []byte{0} //common.ByteSliceToStr(bArr[start : i-1])
	case base.STRING:
		value = []byte{0} // common.ByteSliceToStr(bArr[start : i-1])
	case base.HEX_STRING:
		value = []byte{0}
		// 整型
	case base.HEX:
		value = []byte{0} // string(common.BytesSliceToInt16(bArr[start:length]))
	case base.NUMERAL_TO_INT8:
		value = []byte{0} // string(common.BytesSliceToInt16(bArr[start:length]))
	case base.NUMERAL_TO_INT16:
		value = []byte{0} // string(common.BytesSliceToInt16(bArr[start:length]))
	case base.NUMERAL_TO_INT32:
		value = []byte{0} // string(common.BytesSliceToInt16(bArr[start:length]))
	case base.NUMERAL_TO_INT64:
		value = []byte{0} // string(common.BytesSliceToInt16(bArr[start:length]))

	default:
		value = []byte{0}
	}
	return value
}
