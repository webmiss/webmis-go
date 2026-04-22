package service

/* 数据库 */
type Status struct {
}

/* 公共 */
func (s *Status) Public(name string) map[string]interface{} {
	data := make(map[string]interface{})
	switch name {
	case "role_name":
		data["0"] = "用户"
		data["1"] = "开发"
	case "status_name":
		data["0"] = "禁用"
		data["1"] = "正常"
	}
	return data
}
