package validate

import (
	"common"
)

func Checksum(bytes []byte) (uint16, []byte) {
	l := len(bytes)
	sum := uint16(0)
	for _, b := range bytes[2:l] {
		sum += uint16(b)
	}

	return sum, common.UInt16ToBytesSlice(sum)
}
