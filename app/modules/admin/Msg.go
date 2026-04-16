package admin

import (
	"net/http"
	"webmis/core"
)

/* 消息 */
type Msg struct {
	core.Controller
	partner map[string]map[string]interface{}
}

/* 列表 */
func (c *Msg) List(w http.ResponseWriter, r *http.Request) {
	// 数据
	num := 0
	list := make([]map[string]interface{}, 0)
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0, "data": map[string]interface{}{"num": num, "list": list}})
}
