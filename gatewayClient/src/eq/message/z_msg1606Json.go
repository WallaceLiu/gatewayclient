package message

import (
	"common"
	"encoding/json"
	"fmt"
	"time"
)

type Z_Msg1606Json struct {
	Key            string `json:"key"`
	Bs             string `json:"bs"`             // 标识
	Head_sjbcd     string `json:"head_sjbcd"`     // 数据包长度
	Head_sjbcdjy   string `json:"head_sjbcdjy"`   // 数据包长度校验
	Head_sjbbs     string `json:"head_sjbbs"`     // 数据包ID
	Head_bb        string `json:"head_bb"`        // 协议版本
	Body_mlz       string `json:"body_mlz"`       // 命令字
	Body_sn        string `json:"body_sn"`        // 设备号，字符串
	Body_vin       string `json:"body_vin"`       // 车辆vin，字符串
	Body_pp        string `json:"body_pp"`        // 品牌
	Body_xl        string `json:"body_xl"`        // 系列
	Body_nk        string `json:"body_nk"`        // 年款
	Body_dwxxgs    string `json:"body_dwxxgs"`    // 定位消息个数
	Body_timestamp string `json:"body_timestamp"` // 时间戳，字符串
	Body_cs        string `json:"body_cs"`        // 车速，字符串
	Body_lc        string `json:"body_lc"`        // 里程，字符串
	Body_loc       string `json:"body_loc"`
	Tail_verfiy    string `json:"tail_verfiy"` // 验证码
}

//
func (msg *Z_Msg1606) GetJson() Z_Msg1606Json {
	var msg1606Json Z_Msg1606Json
	msg1606Json.Bs = fmt.Sprintf("%x", msg.bs)
	msg1606Json.Head_sjbcd = fmt.Sprintf("%d", common.BytesSliceToInt16(msg.head_sjbcd))
	msg1606Json.Head_sjbcdjy = fmt.Sprintf("%d", common.BytesSliceToUInt16(msg.head_sjbcdjy))
	msg1606Json.Head_sjbbs = fmt.Sprintf("%d", common.BytesSliceToInt16(msg.head_sjbbs))
	msg1606Json.Head_bb = fmt.Sprintf("%d", common.BytesSliceToInt16(msg.head_bb))
	msg1606Json.Body_mlz = fmt.Sprintf("%x", common.BytesSliceToInt16(msg.body_mlz))
	msg1606Json.Body_sn = common.ByteSliceToStr(msg.body_sn)[0 : len(common.ByteSliceToStr(msg.body_sn))-1]
	msg1606Json.Body_vin = common.ByteSliceToStr(msg.body_vin)[0 : len(common.ByteSliceToStr(msg.body_vin))-1]
	msg1606Json.Body_pp = fmt.Sprintf("%d", common.BytesSliceToInt16(msg.body_pp))
	msg1606Json.Body_xl = fmt.Sprintf("%d", common.BytesSliceToInt16(msg.body_xl))
	msg1606Json.Body_nk = fmt.Sprintf("%d", common.BytesSliceToInt16(msg.body_nk))
	msg1606Json.Body_dwxxgs = fmt.Sprintf("%d", common.BytesSliceToInt16(msg.body_dwxxgs))
	msg1606Json.Body_timestamp = common.ByteSliceToStr(msg.body_timestamp)[0 : len(common.ByteSliceToStr(msg.body_timestamp))-1]
	msg1606Json.Body_cs = common.ByteSliceToStr(msg.body_cs)[0 : len(common.ByteSliceToStr(msg.body_cs))-1]
	msg1606Json.Body_lc = common.ByteSliceToStr(msg.body_lc)[0 : len(common.ByteSliceToStr(msg.body_lc))-1]
	msg1606Json.Body_loc = common.ByteSliceToStr(msg.body_loc)[0 : len(common.ByteSliceToStr(msg.body_loc))-1]
	msg1606Json.Tail_verfiy = fmt.Sprintf("%d", common.BytesSliceToInt16(msg.tail_verfiy))

	msg1606Json.Key = fmt.Sprintf("%d-%s-%s-%d",
		common.GetFnSeq(100000),
		msg1606Json.Body_vin,
		msg1606Json.Body_sn,
		time.Now().UnixNano())

	return msg1606Json
}

//
func (msg *Z_Msg1606Json) ToJson() string {
	b, err := json.Marshal(msg)

	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s", b)
}
