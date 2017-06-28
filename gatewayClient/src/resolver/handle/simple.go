package handle

import (
	"log"
	"resolver/base"
	"resolver/loader"
	"strconv"
)

type Handle struct {
	P_B   []byte         // 二进制消息
	Head  *base.Head     // 消息头
	Msg   *base.Msg      // 消息体
	Reply *base.Msg      // 响应
	Crc   *base.Crc      // crc校验
	R_B   []byte         // 响应
	P     *base.Protocol // 协议
}

// 解析消息
func (handle *Handle) Parser() {

	if len(loader.ProtocolMapBytes) == 0 {
		log.Fatal("Load Protocol Error.")
		return
	}

	handle.P = loader.RecognitionProtocol(handle.P_B, loader.ProtocolMapBytes) // 协议识别

	if handle.P != nil && len(handle.P.Head.Segment) <= 0 {
		log.Fatal("Protocol Or Protocol Head Error, Can not Parse message.")
		return
	}

	firstSegment := handle.P.Head.Segment[0]

	var (
		start  = firstSegment.Start
		funcId string
	)

	log.Println("Parse Start...", len(handle.P_B))

	funcId = parserHead(handle.P_B, handle.P.Head, &start) // 解析消息头

	if len(funcId) == 0 {
		log.Fatal("Found Msg Type Error.")
		return
	}

	handle.Msg = handle.P.Body.FindMsgById(funcId)              //查找消息类型
	handle.Reply = handle.P.Reply.FindMsgById(handle.Msg.Reply) // 查找响应类型

	parserMsg(handle.P_B, handle.Msg, &start) // 解析消息体

	//handle.createReply()
}

// 创建响应
func (handle *Handle) createReply() {
	handle.createBodyHead()
	handle.createBodyReply()
	handle.createBodyCrc()

	var b []byte

	handle.R_B = b
}

// 创建消息
func (handle *Handle) CreateMsg() {
	handle.createBodyHead()
	handle.createBodyMsg()
	handle.createBodyCrc()

	var b []byte

	handle.P_B = b
}

// 解包
// -----------
//
// 解析头
// 十进制消息类型
func parserHead(bArr []byte, head base.Head, start *int) (funcId string) {
	log.Println("Parse Head...")

	parserSegment(bArr, new(base.Msg), head.Segment, start)

	s := base.FindCollectionSegmentFirstById("FuncId", head.Segment)

	if s == nil {
		log.Println("Not Found Message Type.")
		return
	}

	funcId = s.Value

	return funcId
}

// 解析消息体
func parserMsg(bArr []byte, msg *base.Msg, start *int) {
	log.Println("Parse Body...")

	parserSegment(bArr, msg, msg.Segment, start)
}

// 验证CRC
func Crc() {

}

// 解析端-主控
func parserSegment(bArr []byte, msg *base.Msg, collection []base.Segment, start *int) {

	for i, s := range collection {
		if len(s.Segment) <= 0 && !s.IsDel {
			value := transform(bArr, start, &s)
			collection[i].Value = value

			if len(s.SplitId) > 0 {
				split(s.Value, s.SplitId, msg)
			}
		} else {
			if s.IsLoop {
				recursiveSegmentFor(bArr, start, msg, &s)
			} else {
				recursiveSegment(bArr, start, msg, &s)
			}
		}
	}

}

// 解析端-递归
func recursiveSegment(bArr []byte, start *int, msg *base.Msg, segment *base.Segment) {

	if len(bArr) <= *start {
		for i, s := range segment.Segment {
			if len(s.Segment) <= 0 && !s.IsDel {
				value := transform(bArr, start, &s)
				segment.Segment[i].Value = value

				if len(s.SplitId) > 0 {
					split(s.Value, s.SplitId, msg)
				}
			} else {
				if s.IsLoop {
					recursiveSegmentFor(bArr, start, msg, &s)
				} else {
					recursiveSegment(bArr, start, msg, &s)
				}
			}
		}
	}

}

// 解析端-递归
func recursiveSegmentFor(bArr []byte, start *int, msg *base.Msg, segment *base.Segment) {
	if len(bArr) <= *start {
		cntStr := transform(bArr, &segment.Start, segment)

		for _, s := range segment.Segment {
			cnt, _ := strconv.Atoi(cntStr)
			for i := 1; i <= cnt; i++ {
				if len(s.Segment) <= 0 && !s.IsDel {
					value := transform(bArr, start, &s)
					segment.Segment[i].Value = value
					//*start += l

					if len(s.SplitId) > 0 {
						split(s.Value, s.SplitId, msg)
					}
				} else {
					if s.IsLoop {
						recursiveSegmentFor(bArr, start, msg, &s)
					} else {
						recursiveSegment(bArr, start, msg, &s)
					}
				}
			}
		}
	}
}

//
func split(value string, splitId string, msg *base.Msg) {
	seg := base.FindCollectionSegmentFirstById(splitId, msg.Segment)

	for i, s := range seg.Segment {
		if value == s.Id {
			s.Segment[i].IsDel = true
			break
		}
	}
}

//
// 封包
// ----------------
//

// 封装消息头体
func (handle *Handle) createBodyHead() {
	handle.unparserSegment(handle.Head.Segment)
}

// 封装消息体
func (handle *Handle) createBodyMsg() {
	handle.unparserSegment(handle.Msg.Segment)
}

// 封装消息响应体
func (handle *Handle) createBodyReply() {
	handle.unparserSegment(handle.Reply.Segment)
}

// 封装CRC体
func (handle *Handle) createBodyCrc() {

}

// 封包-主控
func (handle *Handle) unparserSegment(collection []base.Segment) {
	for i, s := range collection {
		if len(s.Segment) <= 0 && !s.IsDel {
			collection[i].Byte = unTransform(&s)

			if len(s.SplitId) > 0 {
				split(s.Value, s.SplitId, handle.Msg)
			}
		} else {
			if s.IsLoop {
				handle.unRecursiveSegmentFor(&s)
			} else {
				handle.unRecursiveSegment(&s)
			}
		}
	}
}

// 封包-递归
func (handle *Handle) unRecursiveSegment(segment *base.Segment) {
	for i, s := range segment.Segment {
		if len(s.Segment) <= 0 && !s.IsDel {
			segment.Segment[i].Byte = unTransform(&s)

			if len(s.SplitId) > 0 {
				split(s.Value, s.SplitId, handle.Msg)
			}
		} else {
			if s.IsLoop {
				handle.unRecursiveSegmentFor(&s)
			} else {
				handle.unRecursiveSegment(&s)
			}
		}
	}
}

// 封包-递归
func (handle *Handle) unRecursiveSegmentFor(segment *base.Segment) {
	segment.Byte = unTransform(segment)

	for _, s := range segment.Segment {
		cnt, _ := strconv.Atoi(segment.Value)
		for i := 1; i <= cnt; i++ {
			if len(s.Segment) <= 0 && !s.IsDel {
				segment.Segment[i].Byte = unTransform(&s)

				if len(s.SplitId) > 0 {
					split(s.Value, s.SplitId, handle.Msg)
				}
			} else {
				if s.IsLoop {
					handle.unRecursiveSegmentFor(&s)
				} else {
					handle.unRecursiveSegment(&s)
				}
			}
		}
	}
}
