package base

type Reply struct {
	Id    string `xml:"id,attr"`    // ID
	Title string `xml:"title,attr"` // 标题
	Desc  string `xml:"desc,attr"`  // 描述
	Msg   []Msg  `xml:"msg"`
}

func NewReply(id string, title string, desc string) *Reply {
	return &Reply{
		Id:    id,
		Title: title,
		Desc:  desc}
}

func (reply *Reply) FindMsgById(id string) *Msg {
	for _, msg := range reply.Msg {
		if msg.Id == id {
			return &msg
		}
	}

	return nil
}
