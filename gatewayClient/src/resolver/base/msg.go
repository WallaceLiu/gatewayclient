package base

import ()

type Msg struct {
	Id      string    `xml:"id,attr"`    // ID
	Title   string    `xml:"title,attr"` // 标题
	Desc    string    `xml:"desc,attr"`  // 描述
	Reply   string    `xml:"reply,attr"` // 响应ID
	Segment []Segment `xml:"segment"`
}

func NewMsg(id string, title string, desc string, reply string) *Msg {
	return &Msg{
		Id:    id,
		Title: title,
		Desc:  desc,
		Reply: reply}
}

func (msg *Msg) findMsgSegmentFirstById(id string) *Segment {
	return FindCollectionSegmentFirstById(id, msg.Segment)
}
