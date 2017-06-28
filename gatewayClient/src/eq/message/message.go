package message

func CreateMsg(vin string, sn string) interface{} {
	msg := new(Z_Msg1606).New(vin, sn)
	return msg
}
