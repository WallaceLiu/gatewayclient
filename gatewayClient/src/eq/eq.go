package eq

import (
	"common"
	"config"
	"eq/generator"
	"eq/message"
	"eq/orbit"
	"time"
)

type Eq struct {
	Sn       string
	Vin      string
	Msg      interface{}
	Datetime time.Time
	Lng      float64
	Lat      float64
	Fn       int

	BytesMsg []byte
	StrMsg   string
	JsonMsg  string
}

func NewEq(seq int64) *Eq {
	t, _ := common.GetTimeNow()

	sn := generator.NewSNNumber(seq)
	vin := generator.NewVINNumber(seq)

	return &Eq{
		Sn:       sn,
		Vin:      vin,
		Msg:      message.CreateMsg(vin, sn),
		Datetime: t,
		Lng:      config.LNG,
		Lat:      config.LAT,
		Fn:       common.GetFnSeq(15),
		BytesMsg: []byte{},
		StrMsg:   "",
		JsonMsg:  ""}
}

// 设置设备信息
func (eq *Eq) SetEq() *Eq {
	eq.Lng, eq.Lat = selectOrbitFunc(eq.Lng, eq.Lat, eq.Fn)
	eq.Datetime = eq.Datetime.Add(config.MSG_INTERVAL * time.Second)

	eq.setMsg()

	return eq
}

// 私有方法
// 设置消息的时间、经纬度
func (eq *Eq) setMsg() {
	_, tStr := common.GetTimeNow()

	switch eq.Msg.(type) {
	case message.Z_Msg1606:
		m := eq.Msg.(message.Z_Msg1606)
		m.Set(eq.Lat, eq.Lng, tStr)
		eq.Msg = m
		eq.BytesMsg = m.PackByte()
		eq.StrMsg = m.PackString()
		eq.JsonMsg = m.PackJson()
	}
}

// 私有方法
// 计算经纬度值
// 根据计算公式，计算经纬度值
func selectOrbitFunc(lng float64, lat float64, seq int) (float64, float64) {
	switch seq {
	case 1:
		return orbit.F100_ab01_c0(lng, lat, config.LINEAR_STEP)
	case 2:
		return orbit.F101_ab02_c0(lng, lat, config.LINEAR_STEP)
	case 3:
		return orbit.F102_ab025_c0(lng, lat, config.LINEAR_STEP)
	case 4:
		return orbit.F103_ab03_c0(lng, lat, config.LINEAR_STEP)
	case 5:
		return orbit.F104_ab04_c0(lng, lat, config.LINEAR_STEP)
	case 6:
		return orbit.F105_ab05_c0(lng, lat, config.LINEAR_STEP)
	case 7:
		return orbit.F106_ab06_c0(lng, lat, config.LINEAR_STEP)
	case 8:
		return orbit.F107_ab07_c0(lng, lat, config.LINEAR_STEP)
	case 9:
		return orbit.F108_ab08_c0(lng, lat, config.LINEAR_STEP)
	case 10:
		return orbit.F109_ab09_c0(lng, lat, config.LINEAR_STEP)
	case 11:
		return orbit.F110_ab1_c0(lng, lat, config.LINEAR_STEP)
	case 12:
		return orbit.F111_ab09_c1(lng, lat, config.LINEAR_STEP)
	case 13:
		return orbit.F200_sin(lng, lat, config.TRIGONOMETRIC_STEP)
	case 14:
		return orbit.F201_cos(lng, lat, config.TRIGONOMETRIC_STEP)
	default:
		return orbit.F100_ab01_c0(lng, lat, config.LINEAR_STEP)
	}
}
