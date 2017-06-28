package message

import (
	"log"
	"testing"
)

func Test_CreateMsg(t *testing.T) {
	log.Println(">>> testing...message.go...")

	msg := CreateMsg("VIN0", "SN0")

	log.Println(">>> ", msg)
}
