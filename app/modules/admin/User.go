package admin

import (
	"net/http"
	"webmis/app/config"
	"webmis/app/librarys"
	"webmis/app/models"
	"webmis/app/service"
	"webmis/app/util"
	"webmis/core"
)

/* 用户 */
type User struct {
	core.Controller
}

/* 登录 */
func (c *User) Login(w http.ResponseWriter, r *http.Request) {
	c.Controller.Lang = c.Get(r, "lang")
	// 参数
	json := c.Json(r)
	if json == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	uname := util.Str(c.JsonName(json, "uname"))
	passwd := util.Str(c.JsonName(json, "passwd"))
	vcode := util.Str(c.JsonName(json, "vcode"))
	vcode_url := c.BaseUrl(r, "admin/user/vcode") + "/" + uname + "?" + util.Str(util.Time())
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
		redis := (&core.Redis{}).New("")
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
		redis := (&core.Redis{}).New("")
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
		redis := (&core.Redis{}).New("")
		redis.Set(config.Env().Admin_token_prefix+"_vcode_"+uname, util.Str(util.Time()))
		redis.Expire(config.Env().Admin_token_prefix+"_vcode_"+uname, 24*3600)
		// 返回
		c.GetJSON(w, r, map[string]interface{}{"code": 4000, "msg": c.GetLang("login_verify"), "vcode_url": vcode_url})
		return
	} else {
		// 清除验证码
		redis := (&core.Redis{}).New("")
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
	(&service.TokenAdmin{}).SavePerm(util.Str(data["id"]), util.Str(perm))
	// 登录时间
	ltime := util.Time()
	m = (&models.User{}).New()
	m.Set(map[string]interface{}{"ltime": ltime})
	m.Where("id=?", data["id"])
	m.Update("")
	// Token
	token := (&service.TokenAdmin{}).Create(map[string]interface{}{
		"uid":        data["id"],
		"uname":      uname,
		"name":       data["name"],
		"type":       data["type"],
		"isPasswd":   isPasswd,
		"brand":      data["brand"],
		"shop":       data["shop"],
		"partner":    data["partner"],
		"partner_in": data["partner_in"],
	})
	// 用户信息
	uinfo := map[string]interface{}{
		"uid":        data["id"],
		"uname":      uname,
		"tel":        data["tel"],
		"email":      data["email"],
		"ltime":      util.Date("2006-01-02 15:04:05", ltime),
		"type":       data["type"],
		"nickname":   data["nickname"],
		"department": data["department"],
		"position":   data["position"],
		"name":       data["name"],
		"gender":     data["gender"],
		"birthday":   data["birthday"],
		"img":        (&service.Data{}).Img(util.Str(data["img"]), true),
		"signature":  data["signature"],
	}
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0, "data": map[string]interface{}{"token": token, "uinfo": uinfo, "isPasswd": isPasswd}})
}

/* Token验证 */
func (c *User) Token(w http.ResponseWriter, r *http.Request) {
	c.Controller.Lang = c.Get(r, "lang")
	// 参数
	json := c.Json(r)
	if json == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	token := util.Str(c.JsonName(json, "token"))
	is_uinfo := util.Bool(c.JsonName(json, "uinfo"))
	// 验证
	msg := (&service.TokenAdmin{}).Verify(token, "")
	if msg != "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4001})
		return
	}
	tData := (&service.TokenAdmin{}).Token(token)
	// 用户信息
	uinfo := map[string]interface{}{}
	if is_uinfo {
		m := (&models.User{}).New()
		m.Table("user a")
		m.LeftJoin("user_info AS b", "a.id=b.uid")
		m.Columns(
			"FROM_UNIXTIME(a.ltime) as ltime", "a.tel", "a.email",
			"b.type", "b.nickname", "b.department", "b.position", "b.name", "b.gender", "b.img", "b.signature", "FROM_UNIXTIME(b.birthday, '%Y-%m-%d') as birthday",
		)
		m.Where("a.id=?", tData["uid"])
		uinfo = m.FindFirst("")
		uinfo["uid"] = util.Str(tData["uid"])
		uinfo["uname"] = tData["uname"]
		uinfo["img"] = (&service.Data{}).Img(util.Str(uinfo["img"]), true)
	}
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0, "data": map[string]interface{}{"token_time": tData["time"], "uinfo": uinfo, "isPasswd": tData["isPasswd"]}})
}
