package util

import (
	"fmt"
	"strings"
	"time"
)

/* 时间 */
type TimeType struct{}

/* Time */
func Time() int {
	return int(time.Now().Unix())
}

/* Date: Y-m-d H:i:s */
func Date(format string, timestamp int) string {
	t := time.Now()
	if timestamp != 0 {
		t = time.Unix(int64(timestamp*1000), 0)
	}
	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	hour := t.Hour()
	min := t.Minute()
	sec := t.Second()
	// 替换规则
	replacer := map[string]string{
		"Y": fmt.Sprintf("%04d", year),
		"m": fmt.Sprintf("%02d", month),
		"d": fmt.Sprintf("%02d", day),
		"H": fmt.Sprintf("%02d", hour),
		"i": fmt.Sprintf("%02d", min),
		"s": fmt.Sprintf("%02d", sec),
		"y": fmt.Sprintf("%02d", year%100),
		"n": fmt.Sprintf("%d", month),
		"j": fmt.Sprintf("%d", day),
		"h": fmt.Sprintf("%02d", hour%12),
		"G": fmt.Sprintf("%d", hour),
	}
	// 执行替换
	res := format
	for k, v := range replacer {
		res = strings.ReplaceAll(res, k, v)
	}
	return res
	// if timestamp == 0 {
	// 	timestamp = Time()
	// }
	// str := time.UnixMilli(int64(timestamp) * 1000).Format("2006-01-02 15:04:05")
	// t, _ := time.ParseInLocation(format, str, time.Local)
	// return t.Format(format)
}
