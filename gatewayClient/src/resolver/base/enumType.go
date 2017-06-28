package base

// 字符串常量枚举
const (
	ASCII_COMMA = 44 // 英文逗号 ASCII
	ASCII_END   = 0  // 字符串结束 ASCII
)

// transform 枚举
const (
	ASCII = iota // ASCII

	STRING     // 字符串，0x00结束
	HEX_STRING // 十六进制字符串
	COMMA      // 分隔

	NUMERAL
	NUMERAL_TO_UINT8
	NUMERAL_TO_INT8  // 字节数组转数值
	NUMERAL_TO_UINT16
	NUMERAL_TO_INT16 // 字节数组转数值
	NUMERAL_TO_UINT32
	NUMERAL_TO_INT32 // 字节数组转数值
	NUMERAL_TO_UINT64
	NUMERAL_TO_INT64 // 字节数组转数值

	HEX // 转换成十六进制

	UNKNOWN
)
