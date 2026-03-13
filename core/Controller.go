package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"webmis/app/config/langs"
)

/* 控制器 */
type Controller struct {
	Base
}

/* 获取语言 */
func (c Controller) GetLang(q *http.Request, action string, args ...interface{}) string {
	lang := q.URL.Query().Get("lang")
	if lang == "" {
		lang = "en_US"
	}
	lang = strings.ToLower(lang)
	// 语言包
	var obj interface{}
	if lang == "zh_cn" {
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
func (c Controller) GetJSON(p http.ResponseWriter, q *http.Request, data map[string]interface{}) {
	// Json类型
	p.Header().Set("Content-type", "application/json; charset=utf-8")
	p.WriteHeader(http.StatusOK)
	// 语言
	if data["code"] != nil && data["msg"] == nil {
		data["msg"] = c.GetLang(q, "code_"+fmt.Sprint(data["code"]))
	}
	// 输出
	_ = json.NewEncoder(p).Encode(data)
}

/* Get参数 */
func (Controller) Get(q *http.Request, key string) string {
	return q.URL.Query().Get(key)
}

/* Post参数 */
func (Controller) Post(q *http.Request, key string) string {
	return q.FormValue(key)
}

/* JSON参数 */
func (Controller) Json(q *http.Request) map[string]interface{} {
	if q.Method != http.MethodPost {
		return nil
	}
	var param map[string]interface{}
	err := json.NewDecoder(q.Body).Decode(&param)
	defer q.Body.Close()
	if err != nil {
		return nil
	}
	return param
}
func (Controller) JsonName(param map[string]interface{}, key string) interface{} {
	value, ok := param[key]
	if !ok {
		return nil
	}
	return value
}
