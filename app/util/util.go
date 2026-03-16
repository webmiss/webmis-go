package util

import "strings"

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
