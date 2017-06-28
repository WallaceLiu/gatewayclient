package message

import (
	"bytes"
	"common"
	"config"
	"eq/validate"
	"fmt"
	"time"
)

/*
	位置消息 1606
	--------------------------------------
	格式：
	--------------------------------------
	标识位，2个字节
      	0x55,0xAA

      	消息头：
    	1)，数据包长度，数值，2字节
    	2)，数据包长度校验，数值，2字节
    	3)，数据包ID，数值，1字节
    	4)，协议版本，数值，1字节

	消息体：
        1），命令字，数值，2字节
        2），设备号SN，字符串
        3），车辆VIN，字符串
        4），品牌，数值，1字节
        5），系列，数值，1字节
        6），年款，数值，1字节
        7），定位信息个数，数值，1字节（循环体）
        	1），取得检测数据时间戳，字符串
        	2），定位信息，字符串
        		1），车速，字符串
        		2），当前行程行驶距离，字符串
        		3），经度,纬度,方向,定位时间,定位方式，字符串


      	消息尾：
      	1），验证码，数值，2字节
	--------------------------------------
	示例：
	--------------------------------------
	（145）【消息ID】,
	（5638）【消息类型】,
	（512040090）【设备SN】,
	 (LFPM5ACP4E1A08558)【车辆VIN】,
	（255,255,255）【品牌，系列，年款】,
	（1）【定位消息个数】,
	（1970-01-01 01:42:47）【时间戳】,
	（0,214,E121.535455,N31.279400,0,2015-12-22 10:23:13,2）【定位消息，车速，里程，经度，纬度，方向，定位时间，定位方式】

*/

type Z_Msg1606 struct {
	bs             []byte // 标识
	head_sjbcd     []byte // 数据包长度
	head_sjbcdjy   []byte // 数据包长度校验
	head_sjbbs     []byte // 数据包ID
	head_bb        []byte // 协议版本
	body_mlz       []byte // 命令字
	body_sn        []byte // 设备号，字符串
	body_vin       []byte // 车辆vin，字符串
	body_pp        []byte // 品牌
	body_xl        []byte // 系列
	body_nk        []byte // 年款
	body_dwxxgs    []byte // 定位消息个数
	body_timestamp []byte // 时间戳，字符串
	body_cs        []byte // 车速，字符串
	body_lc        []byte // 里程，字符串
	body_loc       []byte
	tail_verfiy    []byte // 验证码
}

func (msg *Z_Msg1606) New(vin string, sn string) *Z_Msg1606 {
	//var msg Z_Msg1606

	msg.bs = []byte{170, 85}
	msg.head_sjbcd = []byte{0, 0}
	msg.head_sjbcdjy = []byte{0, 0}
	msg.head_sjbbs = []byte{0}
	msg.head_bb = []byte{0}
	msg.body_mlz = []byte{22, 6}
	msg.body_sn = common.GetSn(sn)
	msg.body_vin = common.GetVin(vin)
	msg.body_pp = []byte{255}
	msg.body_xl = []byte{255}
	msg.body_nk = []byte{255}
	msg.body_dwxxgs = []byte{1}

	_, timeStr := common.GetTimeNow()

	msg.body_timestamp = common.StrToByteSlice(timeStr)
	msg.body_cs = common.StrToByteSlice("10")
	msg.body_lc = common.StrToByteSlice("9")
	msg.body_loc = common.GetDefaultLocation()
	msg.tail_verfiy = []byte{0, 0}

	return msg
}

func (msg *Z_Msg1606) Set(lat float64, lng float64, time string) {
	msg.body_timestamp = common.StrToByteSlice(time)
	msg.body_loc = common.GetLocation(lng, lat, 0, time, 1)
}

func (msg *Z_Msg1606) GetBytesForChecksum() []byte {

	msg.head_sjbcd = common.Int16ToBytesSlice(int16(msg.Length()))
	msg.head_sjbcdjy = common.ByteSliceNot(msg.head_sjbcd)

	slice := []byte{}
	slice = append(slice, msg.bs...)           // 2b
	slice = append(slice, msg.head_sjbcd...)   // 2b
	slice = append(slice, msg.head_sjbcdjy...) // 2b
	slice = append(slice, msg.head_sjbbs...)
	slice = append(slice, msg.head_bb...)
	slice = append(slice, msg.body_mlz...)
	slice = append(slice, msg.body_sn...)
	slice = append(slice, msg.body_vin...)
	slice = append(slice, msg.body_pp...)
	slice = append(slice, msg.body_xl...)
	slice = append(slice, msg.body_nk...)
	slice = append(slice, msg.body_dwxxgs...)
	slice = append(slice, msg.body_timestamp...)
	slice = append(slice, msg.body_cs...)
	slice = append(slice, msg.body_lc...)
	slice = append(slice, msg.body_loc...)
	slice = append(slice, msg.tail_verfiy...)

	return slice
}

// 结构体打包成字节数组
func (msg *Z_Msg1606) PackByte() []byte {

	msg.head_sjbcd = common.Int16ToBytesSlice(int16(msg.Length()))
	msg.head_sjbcdjy = common.ByteSliceNot(msg.head_sjbcd)

	slice := []byte{}
	slice = append(slice, msg.bs...)           // 2b
	slice = append(slice, msg.head_sjbcd...)   // 2b
	slice = append(slice, msg.head_sjbcdjy...) // 2b
	slice = append(slice, msg.head_bb...)
	slice = append(slice, msg.head_sjbbs...)
	slice = append(slice, msg.body_mlz...)
	slice = append(slice, msg.body_sn...)
	slice = append(slice, msg.body_vin...)
	slice = append(slice, msg.body_pp...)
	slice = append(slice, msg.body_xl...)
	slice = append(slice, msg.body_nk...)
	slice = append(slice, msg.body_dwxxgs...)
	slice = append(slice, msg.body_timestamp...)
	slice = append(slice, msg.body_cs...)
	slice = append(slice, msg.body_lc...)
	slice = append(slice, msg.body_loc...)

	_, chk := validate.Checksum(slice)

	slice = append(slice, chk...)

	return slice
}

// 结构体打包成字符串
func (msg *Z_Msg1606) PackString() string {

	msg.head_sjbcd = common.Int16ToBytesSlice(int16(msg.Length()))
	msg.head_sjbcdjy = common.ByteSliceNot(msg.head_sjbcd)

	slice := []byte{}
	slice = append(slice, msg.bs...)           // 2b
	slice = append(slice, msg.head_sjbcd...)   // 2b
	slice = append(slice, msg.head_sjbcdjy...) // 2b
	slice = append(slice, msg.head_bb...)
	slice = append(slice, msg.head_sjbbs...)
	slice = append(slice, msg.body_mlz...)
	slice = append(slice, msg.body_sn...)
	slice = append(slice, msg.body_vin...)
	slice = append(slice, msg.body_pp...)
	slice = append(slice, msg.body_xl...)
	slice = append(slice, msg.body_nk...)
	slice = append(slice, msg.body_dwxxgs...)
	slice = append(slice, msg.body_timestamp...)
	slice = append(slice, msg.body_cs...)
	slice = append(slice, msg.body_lc...)
	slice = append(slice, msg.body_loc...)

	_, chk := validate.Checksum(slice)

	msg.tail_verfiy = chk

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d", common.GetFnSeq(100000)%config.REGION_SERVER))
	buf.WriteString(",")
	buf.WriteString(common.ByteSliceToStr(msg.body_sn))
	buf.WriteString(",")
	buf.WriteString(common.ByteSliceToStr(msg.body_vin))
	buf.WriteString(",")
	buf.WriteString(fmt.Sprintf("%d", 1606))
	buf.WriteString(",")
	buf.WriteString(fmt.Sprintf("%d", time.Now().UnixNano()))
	buf.WriteString(",")
	buf.WriteString(fmt.Sprintf("%x", msg.bs))
	buf.WriteString(",")
	buf.WriteString(fmt.Sprintf("%d", common.BytesSliceToInt16(msg.head_sjbcd)))
	buf.WriteString(",")
	buf.WriteString(fmt.Sprintf("%d", common.BytesSliceToUInt16(msg.head_sjbcdjy)))
	buf.WriteString(",")
	buf.WriteString(fmt.Sprintf("%d", common.BytesSliceToInt16(msg.head_sjbbs)))
	buf.WriteString(",")
	buf.WriteString(fmt.Sprintf("%d", common.BytesSliceToInt16(msg.head_bb)))
	buf.WriteString(",")
	buf.WriteString(fmt.Sprintf("%x", common.BytesSliceToInt16(msg.body_mlz)))
	buf.WriteString(",")
	buf.WriteString(common.ByteSliceToStr(msg.body_sn))
	buf.WriteString(",")
	buf.WriteString(common.ByteSliceToStr(msg.body_vin))
	buf.WriteString(",")
	buf.WriteString(fmt.Sprintf("%d", common.BytesSliceToInt16(msg.body_pp)))
	buf.WriteString(",")
	buf.WriteString(fmt.Sprintf("%d", common.BytesSliceToInt16(msg.body_xl)))
	buf.WriteString(",")
	buf.WriteString(fmt.Sprintf("%d", common.BytesSliceToInt16(msg.body_nk)))
	buf.WriteString(",")
	buf.WriteString(fmt.Sprintf("%d", common.BytesSliceToInt16(msg.body_dwxxgs)))
	buf.WriteString(",")
	buf.WriteString(common.ByteSliceToStr(msg.body_timestamp))
	buf.WriteString(",")
	buf.WriteString(common.ByteSliceToStr(msg.body_cs))
	buf.WriteString(",")
	buf.WriteString(common.ByteSliceToStr(msg.body_lc))
	buf.WriteString(",")
	buf.WriteString(common.ByteSliceToStr(msg.body_loc))
	buf.WriteString(",")
	buf.WriteString(fmt.Sprintf("%d", common.BytesSliceToInt16(msg.tail_verfiy)))

	return buf.String()
}

// 结构体打包成JSON
func (msg *Z_Msg1606) PackJson() string {

	msg.head_sjbcd = common.Int16ToBytesSlice(int16(msg.Length()))
	msg.head_sjbcdjy = common.ByteSliceNot(msg.head_sjbcd)

	slice := []byte{}
	slice = append(slice, msg.bs...)           // 2b
	slice = append(slice, msg.head_sjbcd...)   // 2b
	slice = append(slice, msg.head_sjbcdjy...) // 2b
	slice = append(slice, msg.head_bb...)
	slice = append(slice, msg.head_sjbbs...)
	slice = append(slice, msg.body_mlz...)
	slice = append(slice, msg.body_sn...)
	slice = append(slice, msg.body_vin...)
	slice = append(slice, msg.body_pp...)
	slice = append(slice, msg.body_xl...)
	slice = append(slice, msg.body_nk...)
	slice = append(slice, msg.body_dwxxgs...)
	slice = append(slice, msg.body_timestamp...)
	slice = append(slice, msg.body_cs...)
	slice = append(slice, msg.body_lc...)
	slice = append(slice, msg.body_loc...)

	_, chk := validate.Checksum(slice)

	msg.tail_verfiy = chk

	msg1606Json := msg.GetJson()

	return msg1606Json.ToJson()
}

// 消息长度
func (msg *Z_Msg1606) Length() int {
	l := len(msg.head_sjbcd)
	l += len(msg.head_sjbcdjy)
	l += len(msg.head_sjbbs)
	l += len(msg.head_bb)
	l += len(msg.body_mlz)
	l += len(msg.body_sn)
	l += len(msg.body_vin)
	l += len(msg.body_pp)
	l += len(msg.body_xl)
	l += len(msg.body_nk)
	l += len(msg.body_dwxxgs)
	l += len(msg.body_timestamp)
	l += len(msg.body_cs)
	l += len(msg.body_lc)
	l += len(msg.body_loc)
	l += len(msg.tail_verfiy)
	return l
}

// 消息总长度
func (msg *Z_Msg1606) LengthAll() int {
	return msg.Length() + 2
}
