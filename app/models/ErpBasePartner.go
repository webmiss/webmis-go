package models

import (
	"webmis/app/util"
	"webmis/core"
)

/* 系统菜单 */
type ErpBasePartner struct {
	core.Model
}

/* 构造函数 */
func (m *ErpBasePartner) New() *ErpBasePartner {
	m.DBConn("default")
	m.Table("erp_base_partner")
	return m
}

/* 列表 */
func (db *ErpBasePartner) GetList(where []string, columns []string, order string) map[string]map[string]interface{} {
	// 默认
	if columns == nil {
		columns = []string{"name", "status"}
	}
	if order == "" {
		order = "status DESC, sort DESC, name ASC"
	}
	// 字段
	columns = append(columns, "wms_co_id")
	// 查询
	m := db.New()
	m.Columns(columns...)
	m.Where(util.Implode(" AND ", where))
	m.Order(order)
	all := m.Find("")
	// 数据
	data := map[string]map[string]interface{}{}
	for _, v := range all {
		data[v["wms_co_id"].(string)] = v
	}
	return data
}
