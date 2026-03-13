package langs

/* 简体中文 */
type Zh_cn struct {
}

/* Config */
func (Zh_cn) Config(name string) string {
	data := map[string]string{
		// Code
		"code_0":    "成功",
		"code_4000": "参数错误",
		"code_4001": "Token验证失败",
		"code_4010": "暂无数据",
		"code_5000": "服务器错误",
		// Public
		"enable":       "正常",
		"disable":      "禁用",
		"export_limit": "总数%s不能大于%s",
		// Login
		"login_uname":             "请输入用户名/手机/邮箱",
		"login_passwd":            "请输入%s~%s位密码",
		"login_vcode":             "请输入验证码",
		"login_verify":            "帐号或密码错误",
		"login_verify_vcode":      "验证码错误",
		"login_verify_vcode_time": "请%s秒后重试",
		"login_verify_vcode_max":  "超过当天最大上限%s次",
		"login_verify_status":     "该用户已被禁用",
		"login_verify_perm":       "该用户不允许登录",
		// SysUser
		"sys_user_uname":    "请输入用户名/手机/邮箱",
		"sys_user_passwd":   "密码为英文字母开头%s～$s位",
		"sys_user_is_exist": "该用户已存在",
		// SysRole
		"sys_role_name": "角色%s～%s位字符",
		// SysMenus
		"sys_menus_name": "菜单%s～%s位字符",
	}
	return data[name]
}
