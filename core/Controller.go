package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/* 基础控制器 */
type Controller struct {
	// core.Base
}

/* 获取JSON */
func (Controller) GetJSON(c http.ResponseWriter, data map[string]interface{}) {
	c.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(c).Encode(data)
}

/* 输出到控制台 */
func (Controller) Print(content ...interface{}) {
	fmt.Println(content...)
}
