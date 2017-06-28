package util

import (
	"log"
	"testing"
)

func TestConvertTranformEnum(t *testing.T) {
	log.Println(ConvertTranformEnum("Numeral"))
	log.Println(ConvertTranformEnum("NUMERAL_TO_INT16"))
	log.Println(ConvertTranformEnum("string"))
}
