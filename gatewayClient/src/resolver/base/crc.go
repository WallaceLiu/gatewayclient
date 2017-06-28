package base

type Crc struct {
	Id      string    `xml:"id,attr"`    // ID
	Title   string    `xml:"title,attr"` // 标题
	Desc    string    `xml:"desc,attr"`  // 描述
	Segment []Segment `xml:"segment"`
}

func NewCrc(id string, title string, desc string) *Crc {
	return &Crc{
		Id:    id,
		Title: title,
		Desc:  desc}
}

func (crc *Crc) FindCrcSegmentFirstByTitle(id string) *Segment {
	return FindCollectionSegmentFirstById(id, crc.Segment)
}
