package util

import (
	"encoding/json"
	"strings"
)

/* 常用工具 */
type Util struct{}

/* Trim */
func Trim(val string, cutset string) string {
	if cutset == "" {
		return strings.TrimSpace(val)
	} else {
		return strings.Trim(val, cutset)
	}
}

/* Ltrim */
func Ltrim(val string, cutset string) string {
	if cutset == "" {
		return strings.TrimLeft(val, " ")
	} else {
		return strings.TrimLeft(val, cutset)
	}
}

/* Rtrim */
func Rtrim(val string, cutset string) string {
	if cutset == "" {
		return strings.TrimRight(val, " ")
	} else {
		return strings.TrimRight(val, cutset)
	}
}

/* Lower */
func Lower(val string) string {
	return strings.ToLower(val)
}

/* Upper */
func Upper(val string) string {
	return strings.ToUpper(val)
}

/* Explode */
func Explode(val string, sep string) []string {
	return strings.Split(val, sep)
}

/* Implode */
func Implode(val []string, sep string) string {
	return strings.Join(val, sep)
}

/* JsonEncode */
func JsonEncode(data map[string]interface{}) string {
	bytes, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(bytes)
}

/* JsonDecode */
func JsonDecode(jsonStr string) map[string]interface{} {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return data
	}
	return data
}

/* JsonDecodeArr */
func JsonDecodeArr(jsonStr string) []map[string]interface{} {
	var data []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return data
	}
	return data
}
