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
func (m *Index) Index(c http.ResponseWriter, r *http.Request) {
	user := (&models.User{}).New()
	_ = user
	m.GetJSON(c, map[string]interface{}{"code": 200, "data": "Go Api"})
}
