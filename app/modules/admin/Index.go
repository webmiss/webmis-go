package admin

import (
	"net/http"
	"webmis/app/util"
	"webmis/core"
)

/* 控制台 */
type Index struct {
	core.Controller
	partner map[string]map[string]interface{}
}

/* 首页 */
func (c *Index) Index(w http.ResponseWriter, r *http.Request) {
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0, "data": "Go Admin"})
}

/* 软件升级 */
func (c *Index) Version(w http.ResponseWriter, r *http.Request) {
	c.Controller.Lang = c.Get(r, "lang")
	// 参数
	json := c.Json(r)
	if json == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000, "msg": "参数错误!"})
		return
	}
	os := c.JsonName(json, "os").(string)
	local := c.JsonName(json, "version").(string)
	// 验证
	os = util.Lower(os)
	if os != "web" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000, "msg": "[" + os + "]该操作系统不支持更新!"})
		return
	}
	// 数据
	var size = 0
	var version = ""
	var url = ""
	if os == "web" {
		version = "3.0.0"
		url = "https://admin.webmis.vip"
		size = 0
	}
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0, "data": map[string]interface{}{
		"os":      os,
		"version": version,
		"local":   local,
		"size":    size,
		"url":     url,
	}})
}

/* 法定假期 */
func (c *Index) Holiday(w http.ResponseWriter, r *http.Request) {
	c.Controller.Lang = c.Get(r, "lang")
	// 参数
	json := c.Json(r)
	if json == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000, "msg": "参数错误!"})
		return
	}
	date := c.JsonName(json, "date").(string)
	url := "https://go.webmis.vip/upload/img/holiday/"
	// 假期
	holiday := map[string]map[string]interface{}{
		"2026-02-16": map[string]interface{}{"holiday": true, "name": "春节", "img": url + "20260216(360x420).png", "bg": url + "202602(360x50).png"},
		"2026-02-17": map[string]interface{}{"holiday": true, "name": "春节", "img": url + "20260217(360x420).png", "bg": url + "202602(360x50).png"},
		"2026-02-18": map[string]interface{}{"holiday": true, "name": "春节", "img": url + "20260218(360x420).png", "bg": url + "202602(360x50).png"},
		"2026-02-19": map[string]interface{}{"holiday": true, "name": "春节", "img": url + "20260219(360x420).png", "bg": url + "202602(360x50).png"},
		"2026-02-20": map[string]interface{}{"holiday": true, "name": "春节", "img": url + "20260220(360x420).png", "bg": url + "202602(360x50).png"},
		"2026-02-21": map[string]interface{}{"holiday": true, "name": "春节", "img": url + "20260221(360x420).png", "bg": url + "202602(360x50).png"},
		"2026-02-22": map[string]interface{}{"holiday": true, "name": "春节", "img": url + "20260222(360x420).png", "bg": url + "202602(360x50).png"},
		"2026-02-23": map[string]interface{}{"holiday": true, "name": "春节", "img": url + "20260223(360x420).png", "bg": url + "202602(360x50).png"},
		"2026-04-04": map[string]interface{}{"holiday": true, "name": "清明节", "img": "", "bg": ""},
		"2026-04-05": map[string]interface{}{"holiday": true, "name": "清明节", "img": "", "bg": ""},
		"2026-04-06": map[string]interface{}{"holiday": true, "name": "清明节", "img": "", "bg": ""},
		"2026-05-01": map[string]interface{}{"holiday": true, "name": "劳动节", "img": "", "bg": ""},
		"2026-05-02": map[string]interface{}{"holiday": true, "name": "劳动节", "img": "", "bg": ""},
		"2026-05-03": map[string]interface{}{"holiday": true, "name": "劳动节", "img": "", "bg": ""},
		"2026-05-04": map[string]interface{}{"holiday": true, "name": "劳动节", "img": "", "bg": ""},
		"2026-05-05": map[string]interface{}{"holiday": true, "name": "劳动节", "img": "", "bg": ""},
		"2026-06-20": map[string]interface{}{"holiday": true, "name": "端午节", "img": "", "bg": ""},
		"2026-06-21": map[string]interface{}{"holiday": true, "name": "端午节", "img": "", "bg": ""},
		"2026-06-22": map[string]interface{}{"holiday": true, "name": "端午节", "img": "", "bg": ""},
		"2026-09-26": map[string]interface{}{"holiday": true, "name": "中秋节", "img": "", "bg": ""},
		"2026-09-27": map[string]interface{}{"holiday": true, "name": "中秋节", "img": "", "bg": ""},
		"2026-09-28": map[string]interface{}{"holiday": true, "name": "中秋节", "img": "", "bg": ""},
		"2026-10-01": map[string]interface{}{"holiday": true, "name": "国庆节", "img": "", "bg": ""},
		"2026-10-02": map[string]interface{}{"holiday": true, "name": "国庆节", "img": "", "bg": ""},
		"2026-10-03": map[string]interface{}{"holiday": true, "name": "国庆节", "img": "", "bg": ""},
		"2026-10-04": map[string]interface{}{"holiday": true, "name": "国庆节", "img": "", "bg": ""},
		"2026-10-05": map[string]interface{}{"holiday": true, "name": "国庆节", "img": "", "bg": ""},
		"2026-10-06": map[string]interface{}{"holiday": true, "name": "国庆节", "img": "", "bg": ""},
		"2026-10-07": map[string]interface{}{"holiday": true, "name": "国庆节", "img": "", "bg": ""},
	}
	// 返回
	if holiday[date] != nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 0, "data": holiday[date]})
	} else {
		c.GetJSON(w, r, map[string]interface{}{"code": 0, "data": ""})
	}
}
