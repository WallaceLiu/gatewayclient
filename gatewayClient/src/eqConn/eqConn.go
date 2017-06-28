package eqConn

import (
	"common"
	"config"
	"eq"
	"eq/generator"
	"eq/message"
	"net"
)

type EqConn struct {
	num  int
	eq   *eq.Eq   // 设备
	conn net.Conn // TCPIP连接
}

func NewEqConn(seq int64, i int, connect net.Conn) *EqConn {
	t, _ := common.GetTimeNow()

	sn := generator.NewSNNumber(seq)
	vin := generator.NewVINNumber(seq)

	return &EqConn{
		num: i,
		eq: &eq.Eq{
			Sn:       sn,
			Vin:      vin,
			Msg:      message.CreateMsg(vin, sn),
			Datetime: t,
			Lng:      config.LNG,
			Lat:      config.LAT,
			Fn:       common.GetFnSeq(15),
			BytesMsg: []byte{},
			StrMsg:   "",
			JsonMsg:  ""},
		conn: connect}
}
