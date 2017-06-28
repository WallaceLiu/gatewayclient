package eq

import (
	"common"
	"log"
	"testing"
)

// 测试
func Test_NewEq(t *testing.T) {
	log.Println("Testing...New eq...")

	eq := NewEq(common.GetRandFn(int64(1000000)))

	log.Println("SN=", eq.Sn, ",VIN=", eq.Vin, ",lng=", eq.Lng, ",lat=", eq.Lat, ",datetime=", eq.Datetime, ",fn=", eq.Fn)
	log.Println("bytes=", eq.BytesMsg)
}

// 测试
func Test_SetEq(t *testing.T) {
	log.Println("Testing...Set eq...")

	eq := NewEq(common.GetRandFn(int64(1000000)))

	log.Println("SN=", eq.Sn, ",VIN=", eq.Vin, ",datetime=", eq.Datetime, ",fn=", eq.Fn)
	log.Println("lng=", eq.Lng, ",lat=", eq.Lat)

	for i := 1; i <= 100; i++ {
		eq = eq.SetEq()
		log.Println("lng=", eq.Lng, ",lat=", eq.Lat, ",datetime=", eq.Datetime)
	}
}
