package api

import (
	"net/http"
	"webmis/core"
)

/* 接口 */
type Index struct {
	core.Controller
}

/* 首页 */
func (m *Index) Index(c http.ResponseWriter, r *http.Request) {
	m.Print("API")
	m.GetJSON(c, map[string]interface{}{"code": 200, "data": "Go Api"})
}
