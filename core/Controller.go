package core

import (
	"encoding/json"
	"net/http"
	"strings"
)

/* 控制器 */
type Controller struct {
	Base
}

/* 获取JSON */
func (c Controller) GetJSON(p http.ResponseWriter, q *http.Request, data map[string]interface{}) {
	p.Header().Set("Content-type", "application/json; charset=utf-8")
	p.WriteHeader(http.StatusOK)
	lang := q.URL.Query().Get("lang")
	if lang == "" {
		lang = "en_US"
	}
	lang = strings.ToLower(lang)
	// c.Print("lang", lang)
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
