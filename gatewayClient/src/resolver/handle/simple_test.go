package handle

import (
	"log"
	"resolver/loader"
	"testing"
)

var m1608 = []byte{170, 85, 00, 39, 239, 228, 1, 231,
	22, 8,
	53, 49, 50, 48, 52, 48, 50, 55, 51, 0,
	49, 50, 46, 52, 00, 00, 06, 05, 151}

var m1606 = []byte{170, 85, 0, 120, 255,
	135, 0, 4, 22, 6, 48,
	51, 56, 50, 55,
	48, 50, 57, 50, 0,
	76, 70, 80, 77, 53, 65, 67, 80, 54, 71, 49, 65, 51, 51, 54, 53, 55, 0, // VIN
	255, 255, 255, // 品牌系列年款
	0, 1,
	50, 48, 49, 54,
	45, 48, 58, 45, 49, 48, 32,
	48, 48, 58, 48,
	50, 58, 53, 54, 0, 48, 48,
	48, 0, 48, 48, 52, 0, 69,
	49, 49, 55, 46, 51,
	51, 56, 50,
	55, 54, 44, 78, 51, 49,
	46, 55, 57, 55,
	48, 52, 51, 44, 48, 48,
	48, 44, 50, 48,
	49, 54, 45,
	48, 57, 45, 49,
	48, 32, 48,
	48, 58, 48, 50,
	58, 53, 54, 44,
	50, 0, 25, 117}

// 测试消息解析
func TestHandle_Parser1608(t *testing.T) {
	loader.LoadConf()
	log.Println("Protocol Config Files...", len(loader.ProtocolMapBytes))

	var m = new(Handle)
	m.P_B = m1608
	m.Parser()
}

func TestHandle_Parser1606(t *testing.T) {
	loader.LoadConf()
	log.Println("Protocol Config Files...", len(loader.ProtocolMapBytes))

	var m = new(Handle)
	m.P_B = m1606
	m.Parser()
}
