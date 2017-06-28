package orbit

import (
	"config"
	"fmt"
	"math"
	"strconv"
)

// 经纬度
// 东经E+，西经W-，范围：-90~90
// 北纬N+，南纬S-，范围：-180~180

// 线性函数：
// a/b=0.1,c=0
func F100_ab01_c0(v_lng float64, v_lat float64, step float64) (lng float64, lat float64) {
	var x float64 = config.LINEAR_STEP
	var y float64 = 0

	x = x + step
	y = 0.1 * x

	lng, lat = control(v_lng, v_lat, x, y)

	lng, error2 := strconv.ParseFloat(fmt.Sprintf("%.6f", lng), 64)
	if error2 != nil {
	}
	lat, error1 := strconv.ParseFloat(fmt.Sprintf("%.6f", lat), 64)
	if error1 != nil {
	}

	return lng, lat
}

// 线性函数：
// a/b=0.2,c=0
func F101_ab02_c0(v_lng float64, v_lat float64, step float64) (lng float64, lat float64) {
	var x float64 = config.LINEAR_STEP
	var y float64 = 0

	x = x + step
	y = 0.2 * x

	lng, lat = control(v_lng, v_lat, x, y)

	lng, error2 := strconv.ParseFloat(fmt.Sprintf("%.6f", lng), 64)
	if error2 != nil {
	}
	lat, error1 := strconv.ParseFloat(fmt.Sprintf("%.6f", lat), 64)
	if error1 != nil {
	}

	return lng, lat
}

// 线性函数：
// a/b=0.25,c=0
func F102_ab025_c0(v_lng float64, v_lat float64, step float64) (lng float64, lat float64) {
	var x float64 = config.LINEAR_STEP
	var y float64 = 0

	x = x + step
	y = 0.25 * x

	lng, lat = control(v_lng, v_lat, x, y)

	lng, error2 := strconv.ParseFloat(fmt.Sprintf("%.6f", lng), 64)
	if error2 != nil {
	}
	lat, error1 := strconv.ParseFloat(fmt.Sprintf("%.6f", lat), 64)
	if error1 != nil {
	}

	return lng, lat
}

// 线性函数：
// a/b=0.3,c=0
func F103_ab03_c0(v_lng float64, v_lat float64, step float64) (lng float64, lat float64) {
	var x float64 = config.LINEAR_STEP
	var y float64 = 0

	x = x + step
	y = 0.3 * x

	lng, lat = control(v_lng, v_lat, x, y)

	lng, error2 := strconv.ParseFloat(fmt.Sprintf("%.6f", lng), 64)
	if error2 != nil {
	}
	lat, error1 := strconv.ParseFloat(fmt.Sprintf("%.6f", lat), 64)
	if error1 != nil {
	}

	return lng, lat
}

// 线性函数：
// a/b=0.4,c=0
func F104_ab04_c0(v_lng float64, v_lat float64, step float64) (lng float64, lat float64) {
	var x float64 = config.LINEAR_STEP
	var y float64 = 0

	x = x + step
	y = 0.4 * x

	lng, lat = control(v_lng, v_lat, x, y)

	lng, error2 := strconv.ParseFloat(fmt.Sprintf("%.6f", lng), 64)
	if error2 != nil {
	}
	lat, error1 := strconv.ParseFloat(fmt.Sprintf("%.6f", lat), 64)
	if error1 != nil {
	}

	return lng, lat
}

// 线性函数：
// a/b=0.5,c=0
func F105_ab05_c0(v_lng float64, v_lat float64, step float64) (lng float64, lat float64) {
	var x float64 = config.LINEAR_STEP
	var y float64 = 0

	x = x + step
	y = 0.5 * x

	lng, lat = control(v_lng, v_lat, x, y)

	lng, error2 := strconv.ParseFloat(fmt.Sprintf("%.7f", lng), 64)
	if error2 != nil {
	}
	lat, error1 := strconv.ParseFloat(fmt.Sprintf("%.6f", lat), 64)
	if error1 != nil {
	}

	return lng, lat
}

// 线性函数：
// a/b=0.6,c=0
func F106_ab06_c0(v_lng float64, v_lat float64, step float64) (lng float64, lat float64) {
	var x float64 = config.LINEAR_STEP
	var y float64 = 0

	x = x + step
	y = 0.6 * x

	lng, lat = control(v_lng, v_lat, x, y)

	lat, error1 := strconv.ParseFloat(fmt.Sprintf("%.6f", lat), 64)
	if error1 != nil {
	}
	lng, error2 := strconv.ParseFloat(fmt.Sprintf("%.6f", lng), 64)
	if error2 != nil {
	}

	return lng, lat
}

// 线性函数：
// a/b=0.7,c=0
func F107_ab07_c0(v_lng float64, v_lat float64, step float64) (lng float64, lat float64) {
	var x float64 = config.LINEAR_STEP
	var y float64 = 0

	x = x + step
	y = 0.7 * x

	lng, lat = control(v_lng, v_lat, x, y)

	lat, error1 := strconv.ParseFloat(fmt.Sprintf("%.6f", lat), 64)
	if error1 != nil {
	}
	lng, error2 := strconv.ParseFloat(fmt.Sprintf("%.6f", lng), 64)
	if error2 != nil {
	}

	return lng, lat
}

// 线性函数：
// a/b=0.8,c=0
func F108_ab08_c0(v_lng float64, v_lat float64, step float64) (lng float64, lat float64) {
	var x float64 = config.LINEAR_STEP
	var y float64 = 0

	x = x + step
	y = 0.8 * x

	lng, lat = control(v_lng, v_lat, x, y)

	lat, error1 := strconv.ParseFloat(fmt.Sprintf("%.6f", lat), 64)
	if error1 != nil {
	}
	lng, error2 := strconv.ParseFloat(fmt.Sprintf("%.6f", lng), 64)
	if error2 != nil {
	}

	return lng, lat
}

// 线性函数：
// a/b=0.9,c=0
func F109_ab09_c0(v_lng float64, v_lat float64, step float64) (lng float64, lat float64) {
	var x float64 = config.LINEAR_STEP
	var y float64 = 0

	x = x + step
	y = 0.9 * x

	lng, lat = control(v_lng, v_lat, x, y)

	lat, error1 := strconv.ParseFloat(fmt.Sprintf("%.6f", lat), 64)
	if error1 != nil {
	}
	lng, error2 := strconv.ParseFloat(fmt.Sprintf("%.6f", lng), 64)
	if error2 != nil {
	}

	return lng, lat
}

// 线性函数：
// a/b=1,c=0
func F110_ab1_c0(v_lng float64, v_lat float64, step float64) (lng float64, lat float64) {
	var x float64 = config.LINEAR_STEP
	var y float64 = 0

	x = x + step
	y = x

	lng, lat = control(v_lng, v_lat, x, y)

	lat, error1 := strconv.ParseFloat(fmt.Sprintf("%.6f", lat), 64)
	if error1 != nil {
	}
	lng, error2 := strconv.ParseFloat(fmt.Sprintf("%.6f", lng), 64)
	if error2 != nil {
	}
	return lng, lat
}

// 线性函数：
// a/b=0.9,c=1
func F111_ab09_c1(v_lng float64, v_lat float64, step float64) (lng float64, lat float64) {
	var x float64 = config.LINEAR_STEP
	var y float64 = 0

	x = x + step
	y = 0.9*x + config.LINEAR_STEP

	lng, lat = control(v_lng, v_lat, x, y)

	lat, error1 := strconv.ParseFloat(fmt.Sprintf("%.6f", lat), 64)
	if error1 != nil {
	}
	lng, error2 := strconv.ParseFloat(fmt.Sprintf("%.6f", lng), 64)
	if error2 != nil {
	}
	return lng, lat
}

// 正弦函数
// y=sin(x)
func F200_sin(v_lng float64, v_lat float64, step float64) (lng float64, lat float64) {
	var x float64 = 0
	var y float64 = 0

	x = x + step
	y = math.Sin(x)
	lat = v_lat + y
	lng = v_lng + x

	lat, error1 := strconv.ParseFloat(fmt.Sprintf("%.6f", lat), 64)
	if error1 != nil {
	}
	lng, error2 := strconv.ParseFloat(fmt.Sprintf("%.6f", lng), 64)
	if error2 != nil {
	}
	return lng, lat
}

// 余弦函数
// y=cos(x)
func F201_cos(v_lng float64, v_lat float64, step float64) (lng float64, lat float64) {
	var x float64 = 0
	var y float64 = 0

	x = x + step
	y = math.Cos(x)
	lat = v_lat + y
	lng = v_lng + x

	lat, error1 := strconv.ParseFloat(fmt.Sprintf("%.6f", lat), 64)
	if error1 != nil {
	}
	lng, error2 := strconv.ParseFloat(fmt.Sprintf("%.6f", lng), 64)
	if error2 != nil {
	}
	return lng, lat
}

func control(v_lng float64, v_lat float64, x float64, y float64) (lng float64, lat float64) {
	//if v_lng+x >= 90  {
	//	lng = v_lng - x
	//} else {
	//
	//	lng = v_lng + x
	//}
	//
	//if v_lat+y < 180  {
	//	lat = v_lat + y
	//} else {
	//	lat = v_lat - y
	//}

	lng = v_lng + x
	lat = v_lat + y

	if lng >= 90 {
		lng = 0
	}
	if lat > 180 {
		lat = 0
	}

	return lng, lat
}
