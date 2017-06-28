package generator

import (
	"config"
	"fmt"
)

// 创建设备号和车辆底盘编码

// 创建设备号
func NewSNNumber(seq int64) string {
	return config.VIN_PREFIX + fmt.Sprintf("%0"+fmt.Sprintf("%d", config.BIT)+"d", seq)
}

// 创建车辆底盘VIN
func NewVINNumber(seq int64) string {
	return config.SN_PREFIX + fmt.Sprintf("%0"+fmt.Sprintf("%d", config.BIT)+"d", seq)
}
