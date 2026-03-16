package util

import "time"

/* 时间 */
type TimeType struct{}

/* Bool */
func Time() int {
	return int(time.Now().Unix())
}
