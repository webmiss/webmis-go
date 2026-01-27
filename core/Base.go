package core

import "fmt"

/* 基础类 */
type Base struct{}

/* 输出到控制台 */
func (Base) Print(content ...interface{}) {
	fmt.Println(content...)
}
