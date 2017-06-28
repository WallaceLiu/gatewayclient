package base

import (
	"fmt"
	"strconv"
)

type Body struct {
	Id    string `xml:"id,attr"`    // ID
	Title string `xml:"title,attr"` // 标题
	Desc  string `xml:"desc,attr"`  // 描述
	Msg   []Msg  `xml:"msg"`
}

func NewBody(id string, title string, desc string) *Body {
	return &Body{
		Id:    id,
		Title: title,
		Desc:  desc}
}

//
func (body *Body) FindMsgById(idStr string) *Msg {
	id, _ := strconv.Atoi(idStr)
	i := fmt.Sprintf("%x", id)

	for _, msg := range body.Msg {
		if msg.Id == i {
			return &msg
		}
	}

	return nil
}
