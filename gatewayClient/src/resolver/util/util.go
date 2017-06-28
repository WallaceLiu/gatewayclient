package util

import (
	"resolver/base"
	"strings"
)

func ConvertTranformStrTo(str string) int {
	var i int
	switch strings.ToUpper(str) {
	case "ASCII":
		i = base.ASCII

	case "STRING":
		i = base.STRING
	case "HEX_STRING":
		i = base.HEX_STRING

	case "COMMA":
		i = base.COMMA

	case "HEX":
		i = base.HEX
	case "NUMERAL":
		i = base.NUMERAL
	case "NUMERAL_TO_UINT8":
		i=base.NUMERAL_TO_UINT8
	case "NUMERAL_TO_INT8":
		i = base.NUMERAL_TO_INT8
	case "NUMERAL_TO_UINT16":
		i=base.NUMERAL_TO_UINT16
	case "NUMERAL_TO_INT16":
		i = base.NUMERAL_TO_INT16
	case "NUMERAL_TO_UINT32":
		i=base.NUMERAL_TO_UINT32
	case "NUMERAL_TO_INT32":
		i = base.NUMERAL_TO_INT32
	case "NUMERAL_TO_UINT64":
		i= base.NUMERAL_TO_UINT64
	case "NUMERAL_TO_INT64":
		i = base.NUMERAL_TO_INT64

	case "UNKNOWN":
		i = base.UNKNOWN
	}
	return i
}
