package admin

import (
	"net/http"
	"webmis/app/models"
	"webmis/app/service"
	"webmis/app/util"
	"webmis/core"
)

/* 系统菜单 */
type SysMenus struct {
	core.Controller
	menus   map[string][]map[string]interface{}
	permAll map[string]interface{}
}

/* 获取菜单-权限 */
func (c *SysMenus) GetMenusPerm(w http.ResponseWriter, r *http.Request) {
	// 参数
	json := c.Json(r)
	if json == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	token := util.Str(c.JsonName(json, "token"))
	// 验证
	msg := (&service.TokenAdmin{}).Verify(token, "")
	if msg != "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4001})
		return
	}
	// 用户权限
	c.permAll = (&service.TokenAdmin{}).GetPerm(token)
	// 全部菜单
	c._getMenus()
	// 返回
	data := c._getMenusPerm("0")
	c.GetJSON(w, r, map[string]interface{}{"code": 0, "data": data})
}

/* 递归菜单 */
func (c *SysMenus) _getMenusPerm(fid string) []map[string]interface{} {
	data := []map[string]interface{}{}
	M, ok := c.menus[fid]
	if !ok {
		M = data
	}
	for _, val := range M {
		// 菜单权限
		id := util.Str(val["id"])
		perm, ok := c.permAll[id]
		if !ok {
			continue
		}
		// 动作权限
		action := []map[string]interface{}{}
		actionStr := val["action"].(string)
		actionArr := []map[string]interface{}{}
		if actionStr != "" {
			actionArr = util.JsonDecodeArr(actionStr)
		}
		for _, v := range actionArr {
			permVal := util.Int(v["perm"])
			if (util.Int(perm) & permVal) > 0 {
				action = append(action, v)
			}
		}
		// 数据
		value := map[string]interface{}{"url": val["url"], "controller": val["controller"], "action": action}
		langs := map[string]interface{}{"en": val["en_US"], "zh_CN": val["zh_CN"]}
		tmp := map[string]interface{}{"icon": val["ico"], "label": val["title"], "en": val["en"], "value": value, "langs": langs}
		menu := c._getMenusPerm(id)
		if len(menu) > 0 {
			tmp["children"] = menu
		}
		data = append(data, tmp)
	}
	return data
}

/* 全部菜单 */
func (c *SysMenus) _getMenus() {
	m := (&models.SysMenu{}).New()
	m.Columns(
		"id", "fid", "title", "en", "url", "ico", "controller", "sort", "status",
		"en_US", "zh_CN",
		"FROM_UNIXTIME(ctime, '%Y-%m-%d %H:%i:%s') as ctime", "FROM_UNIXTIME(utime, '%Y-%m-%d %H:%i:%s') as utime",
		"action", "remark",
	)
	m.Order("sort, id")
	data := m.Find("")
	c.menus = map[string][]map[string]interface{}{}
	for _, v := range data {
		fid := util.Str(v["fid"])
		if _, ok := c.menus[fid]; !ok {
			c.menus[fid] = []map[string]interface{}{}
		}
		c.menus[fid] = append(c.menus[fid], v)
	}
}
