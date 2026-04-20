package admin

import (
	"net/http"
	"webmis/app/models"
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
	// 统计
	m := (&models.User{}).New()
	m.Table("user as a")
	m.LeftJoin("user_info as b", "a.id=b.uid")
	m.LeftJoin("sys_perm as c", "a.id=c.uid")
	m.LeftJoin("sys_role as d", "c.role=d.id")
	m.Columns("count(*) AS total")
	m.Where(where)
	one := m.FindFirst("")
	// 数据
	total := make(map[string]interface{})
	if one != nil {
		total["total"] = util.Int(one["total"])
	}
	c.Print(one)
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0, "time": util.Date("Y/m/d H:i:s", 0), "data": total})
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
	// 查询
	m := (&models.User{}).New()
	m.Table("user as a")
	m.LeftJoin("user_info as b", "a.id=b.uid")
	m.LeftJoin("sys_perm as c", "a.id=c.uid")
	m.LeftJoin("sys_role as d", "c.role=d.id")
	m.Columns(
		"a.id", "a.uname", "a.email", "a.tel", "a.status", "FROM_UNIXTIME(a.rtime, '%Y-%m-%d %H:%i:%s') as rtime", "FROM_UNIXTIME(a.ltime, '%Y-%m-%d %H:%i:%s') as ltime", "FROM_UNIXTIME(a.utime, '%Y-%m-%d %H:%i:%s') as utime",
		"b.type", "b.nickname", "b.department", "b.position", "b.name", "b.gender", "b.img", "b.remark", "FROM_UNIXTIME(b.birthday, '%Y-%m-%d') as birthday",
		"c.role", "c.perm",
		"d.name as role_name",
	)
	m.Where(where)
	m.Order("a.ltime DESC")
	if order != "" {
		m.Order(order)
	}
	m.Page(page, limit)
	list := m.Find("")
	// 数据
	c.Print(list)
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
	where = append(where, "a.ltime>="+util.Str(start))
	etime, ok := d["etime"]
	if !ok {
		etime = util.Date("Y-m-d", 0)
	}
	end := util.StrToTime(etime.(string) + " 23:59:59")
	where = append(where, "a.ltime<="+util.Str(end))
	// 结果
	return util.Implode(" AND ", where)
}
