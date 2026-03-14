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
func (c *Index) Index(w http.ResponseWriter, r *http.Request) {
	// 查询
	m := (&models.User{}).New()
	m.Columns("id", "uname")
	data := m.Find("")
	// Redis
	rd := (&core.Redis{}).New("default")
	rd.Set("test", "Go Redis")
	c.Print("Data:", data, rd.Get("test"))
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0, "msg": "Go Api"})
}
