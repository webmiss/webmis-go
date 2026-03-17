package admin

import (
	"net/http"
	"webmis/app/config"
	"webmis/app/librarys"
	"webmis/app/models"
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
	// 查询
	m := (&models.User{}).New()
	m.Table("user a")
	m.LeftJoin("user_info AS b", "a.id=b.uid")
	m.LeftJoin("sys_perm AS c", "a.id=c.uid")
	m.LeftJoin("sys_role AS d", "c.role=d.id")
	m.Columns(
		"a.id", "a.status", "a.password", "a.tel", "a.email",
		"b.type", "b.nickname", "b.department", "b.position", "b.name", "b.gender", "FROM_UNIXTIME(b.birthday, '%Y-%m-%d') as birthday", "b.img", "b.signature",
		"c.role", "c.perm", "c.brand", "c.shop", "c.partner", "c.partner_in",
		"d.perm as role_perm",
	)
	m.Where(where)
	data := m.FindFirst("")
	// 是否存在
	if len(data) == 0 {
		// 强制验证码(24小时)
		redis := (&core.Redis{}).New("default")
		redis.Set(config.Env().Admin_token_prefix+"_vcode_"+uname, util.Strval(util.Time()))
		redis.Expire(config.Env().Admin_token_prefix+"_vcode_"+uname, 24*3600)
		// 返回
		c.GetJSON(w, r, map[string]interface{}{"code": 4000, "msg": c.GetLang("login_verify"), "vcode_url": vcode_url})
		return
	} else {
		// 清除验证码
		redis := (&core.Redis{}).New("default")
		redis.Del(config.Env().Admin_token_prefix + "_vcode_" + uname)
	}
	// 是否禁用
	if data["status"] == "0" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000, "msg": c.GetLang("login_verify_status")})
		return
	}
	// 默认密码
	isPasswd := data["password"].(string) == util.Md5(config.Env().Password)
	// 权限
	perm := data["role_perm"]
	if data["perm"] != "" {
		perm = data["perm"]
	}
	if perm == "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000, "msg": c.GetLang("login_verify_perm")})
		return
	}
	c.Print("perm", data, perm, isPasswd)
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0, "data": "Login"})
}
