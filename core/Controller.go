package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"webmis/app/config/langs"
	"webmis/app/util"
)

/* 控制器 */
type Controller struct {
	Base
	Lang string
}

/* 资源地址 */
func (c *Controller) BaseUrl(r *http.Request, url string) string {
	http := "http"
	if r.TLS != nil {
		http = "https"
	}
	return fmt.Sprintf("%s://%s/%s", http, r.Host, url)
}

/* 获取语言 */
func (c *Controller) GetLang(action string, args ...interface{}) string {
	if c.Lang == "" {
		c.Lang = "en_US"
	}
	c.Lang = util.Lower(c.Lang)
	// 语言包
	var obj interface{}
	if c.Lang == "zh_cn" {
		obj = (&langs.Zh_cn{})
	} else {
		obj = (&langs.En_us{})
	}
	// 反射
	class := reflect.ValueOf(obj)
	method := class.MethodByName("Config")
	msg := method.Call([]reflect.Value{reflect.ValueOf(action)})
	if len(msg) == 0 {
		return ""
	}
	return fmt.Sprintf(msg[0].String(), args...)
}

/* 获取JSON */
func (c *Controller) GetJSON(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	// Json类型
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	// 语言
	if data["code"] != nil && data["msg"] == nil {
		data["msg"] = c.GetLang("code_" + util.Strval(data["code"]))
	}
	// 输出
	_ = json.NewEncoder(w).Encode(data)
}

/* Get参数 */
func (*Controller) Get(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

/* Post参数 */
func (*Controller) Post(r *http.Request, key string) string {
	return r.FormValue(key)
}

/* JSON参数 */
func (*Controller) Json(r *http.Request) map[string]interface{} {
	if r.Method != http.MethodPost {
		return nil
	}
	var param map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&param)
	defer r.Body.Close()
	if err != nil {
		return nil
	}
	return param
}
func (*Controller) JsonName(param map[string]interface{}, key string) interface{} {
	value, ok := param[key]
	if !ok {
		return nil
	}
	return value
}
