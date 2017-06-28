package base

type Segment struct {
	Id        string `xml:"id,attr"`        // ID
	Title     string `xml:"title,attr"`     // 标题
	Desc      string `xml:"desc,attr"`      // 描述
	Start     int    `xml:"start,attr"`     // 开始位置
	Len       int    `xml:"len,attr"`       // 长度
	Ty        string `xml:"ty,attr"`        // 数据类型
	IsStop    bool   `xml:"isstop,attr"`    // 是否停止
	IsField   bool   `xml:"isfield,attr"`   // 是否为字段
	IsLoop    bool   `xml:"isloop,attr"`    // 是否为循环
	SplitId   string `xml:"splitid,attr"`   // 拆分
	IsDel     bool   `xml:"isdel,attr"`     // 是否删除
	Transform string `xml:"transform,attr"` // 转换
	Format    string `xml:"fomat,attr"`     // 格式
	Value     string `xml:"value"`
	Byte      []byte
	Segment   []Segment `xml:"segment"` // 集合
}

func NewSegment(id string, title string, desc string,
	start int, len int, ty string,
	isStop bool, isField bool, isLoop bool, splitId string,
	transform string, format string) *Segment {
	return &Segment{
		Id:        id,
		Title:     title,
		Desc:      desc,
		Start:     start,
		Len:       len,
		Ty:        ty,
		SplitId:   splitId,
		IsStop:    isStop,
		IsField:   isField,
		IsLoop:    isLoop,
		Transform: transform,
		Format:    format}
}

func (data *Segment) SetSegment(id string, title string, desc string,
	start int, len int, ty string,
	isStop bool, isField bool, splitId string,
	transform string, format string) {
	data.Id = id
	data.Title = title
	data.Desc = desc
	data.Start = start
	data.Len = len
	data.Ty = ty
	data.IsStop = isStop
	data.IsField = isField
	data.SplitId = splitId
	data.Transform = transform
	data.Format = format
}

func (data *Segment) FindSegmentFirstById(id string) *Segment {
	for _, d := range data.Segment {

		if id == d.Id {
			return &d
		} else {
			if d.Segment != nil && len(d.Segment) > 0 {
				return findSegmentCollectionFirstById(id, d)
			}
		}
	}

	return nil
}
