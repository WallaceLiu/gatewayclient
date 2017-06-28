package base

type Head struct {
	Id      string    `xml:"id,attr"`    // ID
	Title   string    `xml:"title,attr"` // 标题
	Desc    string    `xml:"desc,attr"`  // 描述
	Segment []Segment `xml:"segment"`
}

func NewHead(id string, title string, desc string) *Head {
	return &Head{
		Id:    id,
		Title: title,
		Desc:  desc}
}

func (head *Head) FindHeadSegmentFirstById(id string) *Segment {
	return FindCollectionSegmentFirstById(id, head.Segment)
}
