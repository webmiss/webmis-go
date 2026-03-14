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
func (m *Index) Index(w http.ResponseWriter, r *http.Request) {
	m.GetJSON(w, r, map[string]interface{}{"code": 0, "title": "WebMIS 3.0", "copy": "webmis.vip © 2026"})
}
