package common

import (
	"log"
	"testing"
)

func Test_GetTimeNow(t *testing.T) {
	tT, tStr := GetTimeNow()
	log.Println(tT)
	log.Println(tStr)
}
