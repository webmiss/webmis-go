package admin

import (
	"net/http"
	"webmis/app/util"
	"webmis/core"
)

/* 用户 */
type User struct {
	core.Controller
}

/* 登录 */
func (c *User) Login(w http.ResponseWriter, r *http.Request) {
	// 参数
	json := c.Json(r)
	if json == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000, "msg": "参数错误!"})
		return
	}
	uname := c.JsonName(json, "uname").(string)
	passwd := c.JsonName(json, "passwd").(string)
	vcode := c.JsonName(json, "vcode").(string)
	vcode_url := c.BaseUrl(r, "admin/user/vcode") + "/" + uname + "?" + (&util.Type{}).Strval((&util.TimeType{}).Time())
	c.Print(uname, passwd, vcode, vcode_url)
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0, "data": "Login"})
}
