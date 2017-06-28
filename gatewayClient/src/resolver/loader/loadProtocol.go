package loader

import (
	"encoding/xml"
	"fmt"
	"log"
	"resolver/base"
	"strings"
)

// 协议识别，利用字节数组重新反序列化一个协议结构体
func RecognitionProtocol(bArr []byte, protocolMapBytes map[string]protocolBytes) *base.Protocol {

	var i = 0
	var p *base.Protocol
	for _, protocol := range protocolMapBytes {
		if recognition(bArr, protocol.p) {
			if err := xml.Unmarshal(protocol.bytes, &p); err != nil {
				log.Fatal(err)
			}
			break
		}
		i++
	}
	if i <= len(protocolMapBytes)-1 {
		return p
	} else {
		return new(base.Protocol)
	}
}

// 协议识别
func recognition(bArr []byte, protocol *base.Protocol) bool {
	var (
		i, r = 0, ""
	)

	rArr := strings.Split(protocol.Recognition, ",")

	for i, r = range rArr {
		segment := base.FindCollectionSegmentFirstById(r, protocol.Head.Segment)
		v := strings.ToUpper(fmt.Sprintf("%x", bArr[segment.Start:segment.Len]))

		//log.Printf("RecognitionProtocol:s=%d	l=%d	v=%s	%s",
		//	segment.Start, segment.Len, segment.Value, v)

		if v != segment.Value {
			break
		}
		i++
	}

	if i >= len(rArr)-1 {
		return true
	} else {
		return false
	}
}
