package util

import "time"

/* 时间 */
type TimeType struct{}

/* Bool */
func (t TimeType) Time() int {
	return int(time.Now().Unix())
}
