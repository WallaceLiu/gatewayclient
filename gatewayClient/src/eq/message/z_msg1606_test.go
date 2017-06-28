package message

import (
	"common"
	"config"
	"log"
	"testing"
)

func Test_New1606(t *testing.T) {
	log.Println(">>> testing...msg1606.go...")

	msg := new(Z_Msg1606).New("VIN0", "SN0")
	log.Println(">>> ", msg.PackByte())
}

func Test_Set1606(t *testing.T) {
	log.Println(">>> testing...msg1606.go...")

	_, tStr := common.GetTimeNow()

	msg := new(Z_Msg1606).New("VIN0", "SN0")
	msg.Set(config.LNG, config.LAT, tStr)

	log.Println(">>> ", msg.PackByte())
}

func Test_Pack1606String(t *testing.T) {
	log.Println(">>> testing...msg1606.Parse...")

	_, tStr := common.GetTimeNow()

	msg := new(Z_Msg1606).New("VIN0", "SN0")
	msg.Set(config.LNG, config.LAT, tStr)

	log.Println(">>> ", msg.PackString())
}

func Test_Pack1606JSON(t *testing.T) {
	log.Println(">>> testing...msg1606.Parse...")

	_, tStr := common.GetTimeNow()

	msg := new(Z_Msg1606).New("VIN0", "SN0")
	msg.Set(config.LNG, config.LAT, tStr)

	log.Println(">>> ", msg.PackJson())
}
