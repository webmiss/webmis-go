package home

import (
	"net/http"
	"webmis/core"
)

/* 网站 */
type Index struct {
	core.Controller
}

/* 首页 */
func (m *Index) Index(c http.ResponseWriter, r *http.Request) {
	m.GetJSON(c, map[string]interface{}{"code": 200, "data": "success"})
}
