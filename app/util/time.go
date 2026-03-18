package util

import "time"

/* 时间 */
type TimeType struct{}

/* Time */
func Time() int {
	return int(time.Now().Unix())
}

/* Date: 2006-01-02 15:04:05 */
func Date(format string, timestamp int) string {
	if timestamp == 0 {
		timestamp = Time()
	}
	str := time.UnixMilli(int64(timestamp) * 1000).Format("2006-01-02 15:04:05")
	t, _ := time.ParseInLocation(format, str, time.Local)
	return t.Format("2006-01-02 15:04:05")
}
