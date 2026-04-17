package service

import (
	"webmis/app/config"
	"webmis/app/librarys"
	"webmis/app/models"
	"webmis/app/util"
	"webmis/core"
)

/* Token Admin */
type TokenAdmin struct {
	core.Base
}

/* 验证 */
func (t *TokenAdmin) Verify(token string, urlPerm string) string {
	// Token
	if token == "" {
		return "Token不能为空!"
	}
	tData := (&librarys.Safety{}).Decode(token)
	if tData == nil {
		return "Token验证失败!"
	}
	// 是否过期
	uid := util.Str(tData["uid"])
	key := config.Env().Admin_token_prefix + "_token_" + uid
	redis := (&core.Redis{}).New("")
	time := redis.Ttl(key)
	if time < 1 {
		return "请重新登录!"
	}
	// 单点登录
	access_token := redis.Get(key)
	if config.Env().Admin_token_sso && util.Md5(token) != access_token {
		return "强制退出!"
	}
	// 是否续期
	if config.Env().Admin_token_auto {
		redis.Expire(key, config.Env().Admin_token_time)
		redis.Expire(config.Env().Admin_token_prefix+"_perm_"+uid, config.Env().Admin_token_time)
	}
	// URL权限
	if urlPerm == "" {
		return ""
	}
	arr := util.Explode("/", urlPerm)
	action := util.Explode("?", arr[len(arr)-1])[0]
	arr = arr[:len(arr)-1]
	controller := util.Implode("/", arr)
	// 查询菜单
	m := (&models.SysMenu{}).New()
	m.Columns("id", "action")
	m.Where("controller=?", controller)
	data := m.FindFirst("")
	if len(data) == 0 {
		return "菜单验证无效!"
	}
	// 验证菜单
	id := util.Str(data["id"])
	perm := t.GetPerm(token)
	if _, ok := perm[id]; !ok {
		return "无权访问菜单!"
	}
	// 验证动作
	permVal := 0
	actionVal := util.Int(perm[id])
	permArr := util.JsonDecodeArr(util.Str(data["action"]))
	for _, v := range permArr {
		if action == v["action"].(string) {
			permVal = util.Int(v["perm"])
			break
		}
	}
	if (actionVal & permVal) == 0 {
		return "无权访问动作!"
	}
	return ""
}

/* 权限-保存 */
func (t *TokenAdmin) SavePerm(uid string, perm string) {
	key := config.Env().Admin_token_prefix + "_perm_" + uid
	redis := (&core.Redis{}).New("")
	redis.Set(key, perm)
	redis.Expire(key, config.Env().Admin_token_time)
}

/* 权限-获取 */
func (t *TokenAdmin) GetPerm(token string) map[string]interface{} {
	arr := map[string]interface{}{}
	// Token
	if token == "" {
		return arr
	}
	tData := (&librarys.Safety{}).Decode(token)
	if tData == nil {
		return arr
	}
	// 权限
	uid := util.Str(tData["uid"])
	redis := (&core.Redis{}).New("")
	permStr := redis.Get(config.Env().Admin_token_prefix + "_perm_" + uid)
	if permStr == "" {
		return arr
	}
	// 拆分
	perm := util.Explode(" ", permStr)
	for _, v := range perm {
		tmp := util.Explode(":", v)
		arr[tmp[0]] = tmp[1]
	}
	return arr
}

/* 生成 */
func (t *TokenAdmin) Create(data map[string]interface{}) string {
	// 登录时间
	data["l_time"] = util.Date("2006-01-02 15:04:05", 0)
	token := (&librarys.Safety{}).Encode(data)
	// 缓存Token
	key := config.Env().Admin_token_prefix + "_token_" + util.Str(data["uid"])
	redis := (&core.Redis{}).New("")
	redis.Set(key, util.Md5(token))
	redis.Expire(key, config.Env().Admin_token_time)
	return token
}

/* 解析 */
func (t *TokenAdmin) Token(token string) map[string]interface{} {
	data := (&librarys.Safety{}).Decode(token)
	if data == nil {
		return nil
	}
	// 过期时间
	redis := (&core.Redis{}).New("")
	data["time"] = redis.Ttl(config.Env().Admin_token_prefix + "_token_" + util.Str(data["uid"]))
	return data
}
