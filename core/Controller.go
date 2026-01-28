package core

import (
	"encoding/json"
	"io"
	"net/http"
)

/* 基础控制器 */
type Controller struct {
	Base
}

/* 获取JSON */
func (Controller) GetJSON(c http.ResponseWriter, data map[string]interface{}) {
	c.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(c).Encode(data)
}

/* Get参数 */
func (Controller) Get(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

/* Post参数 */
func (Controller) Post(r *http.Request, key string) string {
	return r.FormValue(key)
}

/* JSON参数 */
func (Controller) Json(r *http.Request) map[string]interface{} {
	if r.Header.Get("Content-Type") != "application/json" {
		return nil
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil
	}
	param := map[string]interface{}{}
	err = json.Unmarshal(body, &param)
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
