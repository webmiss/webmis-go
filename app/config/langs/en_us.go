package langs

/* English */
type En_us struct {
}

/* Config */
func (En_us) Config(name string) string {
	data := map[string]string{
		// Code
		"code_0":    "Success",
		"code_4000": "Parameter error",
		"code_4001": "Token verification failed",
		"code_4010": "No data available",
		"code_5000": "Server error",
		// Public
		"enable":       "Enable",
		"disable":      "Disable",
		"export_limit": "The total number cannot exceed %s",
		// Login
		"login_uname":             "Uname \\ PhoneNumber \\ Email",
		"login_passwd":            "Password %s-%s characters",
		"login_vcode":             "Please enter the verification code",
		"login_verify":            "Account or password error",
		"login_verify_vcode":      "Verification code error",
		"login_verify_vcode_time": "Please try again in %s seconds",
		"login_verify_vcode_max":  "Exceeding the maximum limit of %s times on the same day",
		"login_verify_status":     "This user has been disabled",
		"login_verify_perm":       "This user is not allowed to login",
		// SysUser
		"sys_user_uname":    "Uname \\ PhoneNumber \\ Email",
		"sys_user_passwd":   "Password %s-%s characters",
		"sys_user_is_exist": "The user already exists",
		// SysRole
		"sys_role_name": "Role %s-%s characters",
		// SysMenus
		"sys_menus_name": "Menus %s-%s characters",
	}
	return data[name]
}
