package config

import (
	"encoding/json"
	io "io/ioutil"
	"log"
)

/*
	配置
	MAX_CONNECT - 最大连接数
	VIN_PREFIX - 车辆底盘VIN前缀
	SN_PREFIX - 设备号前缀
	BIT - 设备号SN和车辆底盘VIN码总位数

	MIN    - 最小连接序号
	MAX    - 最大连接序号

	LAT，LNG - 经度，纬度，默认为雍和宫 （116.424032，39.9535）

	LINEAR_STEP - 线性函数步长
	TRIGONOMETRIC_STEP - 三角函数步长

	MSG_INTERVAL - 时间间隔

	TIMEOUT - TCP/IP连接超时时间
*/
const (
	MAX_CONNECT = 2

	VIN_PREFIX = "VIN"
	SN_PREFIX  = "SN"
	BIT        = 6

	MIN = 0
	MAX = MAX_CONNECT + MIN

	LNG                = 39.9535
	LAT                = 116.424032
	LINEAR_STEP        = 0.000001
	TRIGONOMETRIC_STEP = 1

	MSG_INTERVAL = 1

	TIMEOUT = 2
)

const (
	DB_TYPE    = "mysql"
	SK_DB_URL  = "root:123456@tcp(10.1.8.166:3306)/sk?autocommit=true"       // 斯柯达
	VW_DB_URL  = "root:123456@tcp(10.1.8.166:3306)/vw?autocommit=true"       // 上汽大众
	MSG_DB_URL = "root:123456@tcp(10.1.8.166:3306)/scsj_msg?autocommit=true" // obd 消息
)

const (
	REDIS_HOST = "10.1.8.185"
	REDIS_PORT = 6379
)

const (
	GATEWAY_IPANDHOST = "192.168.5.22:11000"
	GATEWAY_IP        = "192.168.5.22"
	GATEWAY_HOST      = "11000"
)

const (
	MAX_CONNECT_TO_KAKA = 200
	MIN_TO_KAKA         = 0
	MAX_TO_KAKA         = MAX_CONNECT_TO_KAKA + MIN_TO_KAKA
	KAFKA_ADDRS         = "10.1.12.108:9092,10.1.12.109:9092,10.1.12.110:9092"
	KAFKA_TOPIC         = "obd2"
)

const (
	REGION_SERVER = 8
)

const (
	FOR_HBASE = "1"
)

func GetConfig(filename string, v interface{}) {
	if b, err := io.ReadFile(filename); err != nil {
		log.Println(err)
		return
	} else {
		if err = json.Unmarshal(b, &v); err != nil {
			log.Println(err)
			return
		}
	}
}
