package main

import (
	"simplelotus/lotuslib"
)

const (
	ip   = "0.0.0.0"
	port = 8888
)

func main() {
	tcplotus.TcpLotusMain(ip, port)
}
