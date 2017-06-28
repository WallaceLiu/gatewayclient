package base

import (
	"encoding/json"
	"fmt"
)

// 协议
type Protocol struct {
	Title       string `xml:"title,attr"`       // 标题
	Recognition string `xml:"recognition,attr"` // 描述
	Head        Head   `xml:"head"`
	Body        Body   `xml:"body"`
	Reply       Reply  `xml:"reply"`
	Crc         Crc    `xml:"crc"`
}

func NewProtocol() *Protocol {
	return &Protocol{
		Title:       "",
		Recognition: ""}
}

func (protocol *Protocol) ToJson() string {
	b, err := json.Marshal(protocol)

	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s", b)
}
