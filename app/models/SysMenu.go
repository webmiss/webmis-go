package models

import "webmis/core"

/* 系统菜单 */
type SysMenu struct {
	core.Model
}

/* 构造函数 */
func (m *SysMenu) New() *SysMenu {
	m.DBConfig("default")
	m.Table("sys_menus")
	return m
}
