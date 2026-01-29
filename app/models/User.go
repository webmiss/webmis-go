package models

import "webmis/core"

/* 用户模型 */
type User struct {
	core.Model
}

/* 构造函数 */
func (m *User) New() *User {
	m.DBConn("default")
	return m
}
