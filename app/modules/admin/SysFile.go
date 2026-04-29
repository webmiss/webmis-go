package admin

import (
	"net/http"
	"webmis/app/config"
	"webmis/app/librarys"
	"webmis/app/service"
	"webmis/app/util"
	"webmis/core"
)

var dirRoot = "upload/"

/* 文件管理 */
type SysFile struct {
	core.Controller
}

/* 列表 */
func (c *SysFile) List(w http.ResponseWriter, r *http.Request) {
	c.Controller.Lang = c.Get(r, "lang")
	// 参数
	json := c.Json(r)
	if json == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	token := util.Str(c.JsonName(json, "token"))
	path := util.Str(c.JsonName(json, "path"))
	// 验证
	msg := (&service.TokenAdmin{}).Verify(token, r.RequestURI)
	if msg != "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4001})
		return
	}
	if path == "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	// 数据
	(&librarys.FileEo{}).New(config.Env().RootDir + dirRoot)
	list := (&librarys.FileEo{}).List(path)
	// 返回
	c.GetJSON(w, r, map[string]interface{}{
		"code": 0,
		"time": util.Date("Y/m/d H:i:s", 0),
		"data": map[string]interface{}{"url": c.BaseUrl(r, dirRoot), "list": list},
	})
}

/* 新建文件夹 */
func (c *SysFile) Mkdir(w http.ResponseWriter, r *http.Request) {
	c.Controller.Lang = c.Get(r, "lang")
	// 参数
	json := c.Json(r)
	if json == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	token := util.Str(c.JsonName(json, "token"))
	path := util.Str(c.JsonName(json, "path"))
	name := util.Str(c.JsonName(json, "name"))
	// 验证
	msg := (&service.TokenAdmin{}).Verify(token, r.RequestURI)
	if msg != "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4001})
		return
	}
	if path == "" || name == "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	// 数据
	(&librarys.FileEo{}).New(config.Env().RootDir + dirRoot)
	res := (&librarys.FileEo{}).Mkdir(path + name)
	if !res {
		c.GetJSON(w, r, map[string]interface{}{"code": 5000})
		return
	}
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0})
}

/* 重命名 */
func (c *SysFile) Rename(w http.ResponseWriter, r *http.Request) {
	c.Controller.Lang = c.Get(r, "lang")
	// 参数
	json := c.Json(r)
	if json == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	token := util.Str(c.JsonName(json, "token"))
	path := util.Str(c.JsonName(json, "path"))
	name := util.Str(c.JsonName(json, "name"))
	rename := util.Str(c.JsonName(json, "rename"))
	// 验证
	msg := (&service.TokenAdmin{}).Verify(token, r.RequestURI)
	if msg != "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4001})
		return
	}
	if path == "" || name == "" || rename == "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	// 数据
	(&librarys.FileEo{}).New(config.Env().RootDir + dirRoot)
	res := (&librarys.FileEo{}).Rename(path+rename, path+name)
	if !res {
		c.GetJSON(w, r, map[string]interface{}{"code": 5000})
		return
	}
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0})
}

/* 删除 */
func (c *SysFile) Remove(w http.ResponseWriter, r *http.Request) {
	c.Controller.Lang = c.Get(r, "lang")
	// 参数
	json := c.Json(r)
	if json == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	token := util.Str(c.JsonName(json, "token"))
	path := util.Str(c.JsonName(json, "path"))
	data := c.JsonName(json, "data").([]interface{})
	// 验证
	msg := (&service.TokenAdmin{}).Verify(token, r.RequestURI)
	if msg != "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4001})
		return
	}
	if path == "" || data == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	// 数据
	(&librarys.FileEo{}).New(config.Env().RootDir + dirRoot)
	for _, v := range data {
		(&librarys.FileEo{}).RemoveAll(path + util.Str(v))
	}
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0})
}

/* 上传 */
func (c *SysFile) Upload(w http.ResponseWriter, r *http.Request) {
	c.Controller.Lang = c.Get(r, "lang")
	// 参数
	json := c.Json(r)
	if json == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	token := util.Str(c.JsonName(json, "token"))
	path := util.Str(c.JsonName(json, "path"))
	// 验证
	msg := (&service.TokenAdmin{}).Verify(token, r.RequestURI)
	if msg != "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4001})
		return
	}
	if path == "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	// 数据
	_, fileHeader, _ := r.FormFile("file")
	img := (&librarys.Upload{}).File(fileHeader, map[string]interface{}{"path": dirRoot + path, "bind": nil})
	if img == "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 5000})
		return
	}
	// 返回
	c.GetJSON(w, r, map[string]interface{}{"code": 0})
}

/* 下载 */
func (c *SysFile) Down(w http.ResponseWriter, r *http.Request) {
	c.Controller.Lang = c.Get(r, "lang")
	// 参数
	json := c.Json(r)
	if json == nil {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	token := util.Str(c.JsonName(json, "token"))
	path := util.Str(c.JsonName(json, "path"))
	filename := util.Str(c.JsonName(json, "filename"))
	// 验证
	msg := (&service.TokenAdmin{}).Verify(token, r.RequestURI)
	if msg != "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4001})
		return
	}
	if path == "" || filename == "" {
		c.GetJSON(w, r, map[string]interface{}{"code": 4000})
		return
	}
	// 数据
	(&librarys.FileEo{}).New(config.Env().RootDir + dirRoot)
	data := (&librarys.FileEo{}).Bytes(path + filename)
	// 返回
	c.GetFile(w, data, map[string]string{"Content-Type": "application/octet-stream"})
}
