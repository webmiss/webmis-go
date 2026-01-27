package core

import (
	"encoding/json"
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
