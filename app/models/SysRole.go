package models

import "webmis/core"

/* 角色 */
type SysRole struct {
	core.Model
}

/* 构造函数 */
func (m *SysRole) New() *SysRole {
	m.DBConn("default")
	m.Table("sys_role")
	return m
}
