package web

import (
	"net/http"
	"webmis/core"
)

/* 网站 */
type Index struct {
	core.Controller
}

/* 首页 */
func (m *Index) Index(p http.ResponseWriter, q *http.Request) {
	m.GetJSON(p, q, map[string]interface{}{"code": 0, "title": "WebMIS 3.0", "copy": "webmis.vip © 2026"})
}
