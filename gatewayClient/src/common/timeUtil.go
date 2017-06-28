package common

import (
	"fmt"
	"time"
)

// 获得当前日期函数
func GetTimeNow() (t time.Time, tStr string) {
	t = time.Now()
	year, month, day := t.Local().Date()
	h, m, s := t.Local().Clock()
	tStr = fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d", year, month, day, h, m, s)

	return t, tStr
}
