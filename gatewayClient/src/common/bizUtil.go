package common

import (
	"config"
	"fmt"
	"math/rand"
	"time"
)

// 获得随机函数编号
func GetRandFn(fn int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(fn)
}

func GetFnSeq(fn int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(fn)
}

func GetRandomInt(i int) int {
	return GetFnSeq(i)
}

// 获得随机函数编号
func GetRandCode() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(config.MAX_CONNECT)
}

// 默认设备SN号
func GetSn(sn string) []byte {
	if config.FOR_HBASE == "1" {
		return StrToByteSliceNoB0(sn)
	}

	return StrToByteSlice(sn)
}

// 默认车辆底盘VIN码
func GetVin(vin string) []byte {
	if config.FOR_HBASE == "1" {
		return StrToByteSliceNoB0(vin)
	}
	return StrToByteSlice(vin)
}

//
func GetDefaultLocation() []byte {
	_, tStr := GetTimeNow()
	return GetLocation(config.LNG, config.LAT, 0, tStr, 1)
}

//
func GetLocation(lng float64, lat float64, dir int, dateTime string, way int) []byte {
	loc := fmt.Sprintf("E%f,N%f,%d,%s,%d", lng, lat, dir, dateTime, way)
	return StrToByteSlice(loc)
}
