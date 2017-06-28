package loader

import (
	"log"
	"testing"
)

func Test_RecognitionProtocol(t *testing.T) {

	LoadConf()

	bArr := []byte{85, 170}
	p := RecognitionProtocol(bArr, ProtocolMapBytes)
	log.Println(p.Title, bArr)

	bArr = []byte{64, 64}
	p = RecognitionProtocol(bArr, ProtocolMapBytes)
	log.Println(p.Title, bArr)

	bArr = []byte{1, 1}
	p = RecognitionProtocol(bArr, ProtocolMapBytes)
	log.Println(p.Title, bArr)
}
