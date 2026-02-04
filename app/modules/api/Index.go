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
func (c *Index) Index(p http.ResponseWriter, q *http.Request) {
	// 查询
	m := (&models.User{}).New()
	m.Columns("id", "uname")
	data := m.Find("")
	// Redis
	r := (&core.Redis{}).New("default")
	r.Set("test", "Go Redis")
	c.Print("Data:", data, r.Get("test"))
	// 返回
	c.GetJSON(p, map[string]interface{}{"code": 200, "msg": "Go Api"})
}
