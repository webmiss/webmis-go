package admin

import (
	"net/http"
	"webmis/app/config"
	"webmis/app/librarys"
	"webmis/app/util"
	"webmis/core"
)

/* 用户 */
type User struct {
	core.Controller
}

/* 登录 */
func (c *User) Login(w http.ResponseWriter, r *http.Request) {
	c.Controller.Lang = r.URL.Query().Get("lang")
	// 参数
	json := c.Json(r)
	if json == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000, "msg": "参数错误!"})
		return
	}
	uname := c.JsonName(json, "uname").(string)
	passwd := c.JsonName(json, "passwd").(string)
	vcode := c.JsonName(json, "vcode").(string)
	vcode_url := c.BaseUrl(r, "admin/user/vcode") + "/" + uname + "?" + util.Strval(util.Time())
	// 验证用户名
	if !(&librarys.Safety{}).IsRight("uname", uname) && !(&librarys.Safety{}).IsRight("tel", uname) && !(&librarys.Safety{}).IsRight("email", uname) {
		c.GetJSON(w, r, map[string]interface{}{"code": 4001, "msg": c.GetLang("login_uname")})
		return
	}
	// 登录方式
	where := ""
	vcode = util.Lower(util.Trim(vcode, ""))
	if passwd != "" {
		// 密码长度
		if !(&librarys.Safety{}).IsRight("passwd", passwd) {
			c.GetJSON(w, r, map[string]interface{}{"code": 4000, "msg": c.GetLang("login_passwd")})
			return
		}
		// 验证码
		redis := (&core.Redis{}).New("default")
		code := redis.Get(config.Env().Admin_token_prefix + "_vcode_" + uname)
		if code != "" {
			if len(code) != 4 {
				c.GetJSON(w, r, map[string]interface{}{"code": 4001, "msg": c.GetLang("login_vcode"), "vcode_url": vcode_url})
				return
			} else if code != vcode {
				c.GetJSON(w, r, map[string]interface{}{"code": 4002, "msg": c.GetLang("login_verify_vcode"), "vcode_url": vcode_url})
				return
			}
		}
		where = "(a.uname='" + uname + "' OR a.tel='" + uname + "' OR a.email='" + uname + "') AND a.password='" + util.Md5(passwd) + "'"
		c.Print("redis", config.Env().Admin_token_prefix, code, where)
	} else {
		// 验证码
		redis := (&core.Redis{}).New("default")
		code := redis.Get(config.Env().Admin_token_prefix + "_vcode_" + uname)
		if code != "" || code != vcode {
			c.GetJSON(w, r, map[string]interface{}{"code": 4000, "msg": c.GetLang("login_verify_vcode")})
			return
		}
		where = "a.tel='" + uname + "'"
	}
	c.Print(uname, passwd, vcode, where)
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0, "data": "Login"})
}
