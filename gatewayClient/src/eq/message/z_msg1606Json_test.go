package message

import (
	"common"
	"config"
	"fmt"
	"log"
	"testing"
)

func Test_GetMsg1606Json(t *testing.T) {
	fmt.Println(">>> testing...msg1606Json.GetMsg1606Json...")

	_, tStr := common.GetTimeNow()

	msg := new(Z_Msg1606).New("VIN0", "SN0")
	msg.Set(config.LNG, config.LAT, tStr)
	msgJson := msg.GetJson()

	log.Println(">>> ", msgJson.ToJson())
}
