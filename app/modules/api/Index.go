package api

import (
	"net/http"
	"webmis/app/models"
	"webmis/core"
)

/* 接口 */
type Index struct {
	core.Controller
}

/* 首页 */
func (c *Index) Index(p http.ResponseWriter, r *http.Request) {
	m := (&models.User{}).New()
	m.Columns("id")
	m.Where("name=?", "admin")
	m.FindFirst("", nil)
	// sql, _ := m.SelectSQL()
	//
	// c.Print(sql)
	c.GetJSON(p, map[string]interface{}{"code": 200, "data": "Go Api"})
}
