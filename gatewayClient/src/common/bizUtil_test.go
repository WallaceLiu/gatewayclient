package common

import (
	"log"
	"testing"
)

func Test_GetLocation(t *testing.T) {
	s := GetLocation(1.0, 2.0, 1, "2016-12-06 11:11:00", 1)
	log.Println(s)
}

func Test_GetDefaultSn(t *testing.T) {
	s := GetSn("1")
	log.Println(s)
}
