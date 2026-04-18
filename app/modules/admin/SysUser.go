package admin

import (
	"net/http"
	"webmis/app/service"
	"webmis/app/util"
	"webmis/core"
)

/* 系统用户 */
type SysUser struct {
	core.Controller
}

/* 统计 */
func (c *SysUser) Total(w http.ResponseWriter, r *http.Request) {
	c.Controller.Lang = c.Get(r, "lang")
	// 参数
	json := c.Json(r)
	if json == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	token := util.Str(c.JsonName(json, "token"))
	data := c.JsonName(json, "data").(map[string]interface{})
	// 验证
	msg := (&service.TokenAdmin{}).Verify(token, "")
	if msg != "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4001})
		return
	}
	if data == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	// 条件
	where := c.getWhere(data)
	c.Print(where)
	// 数据
	list := map[string]interface{}{}
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0, "time": util.Date("Y/m/d H:i:s", 0), "data": list})
}

/* 列表 */
func (c *SysUser) List(w http.ResponseWriter, r *http.Request) {
	c.Controller.Lang = c.Get(r, "lang")
	// 参数
	json := c.Json(r)
	if json == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	token := util.Str(c.JsonName(json, "token"))
	data := c.JsonName(json, "data").(map[string]interface{})
	page := util.Int(c.JsonName(json, "page"))
	limit := util.Int(c.JsonName(json, "limit"))
	order := util.Str(c.JsonName(json, "order"))
	// 验证
	msg := (&service.TokenAdmin{}).Verify(token, r.RequestURI)
	if msg != "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4001})
		return
	}
	if data == nil || page == 0 || limit == 0 {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	// 条件
	where := c.getWhere(data)
	c.Print(where, order)
	// 数据
	list := map[string]interface{}{}
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0, "time": util.Date("Y/m/d H:i:s", 0), "data": list})
}

/* 搜索条件 */
func (c *SysUser) getWhere(d map[string]interface{}) string {
	where := []string{}
	// 时间
	stime, ok := d["stime"]
	if !ok {
		stime = util.Date("Y-m-d", 0)
	}
	start := util.StrToTime(stime.(string) + " 00:00:00")
	c.Print(stime.(string) + " 00:00:00")
	where = append(where, "stime>="+util.Str(start))
	return util.Implode(" AND ", where)
}
