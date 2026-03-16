package config

type EnvType struct {
	Mode                  string
	ServerHost            string
	ServerPort            string
	Key                   string
	Password              string
	Img_url               string
	Admin_token_prefix    string
	Admin_token_time      int
	Admin_token_auto      bool
	Admin_token_sso       bool
	Api_token_prefix      string
	Api_token_time        int
	Api_token_auto        bool
	Api_token_sso         bool
	Supplier_token_prefix string
	Supplier_token_time   int
	Supplier_token_auto   bool
	Supplier_token_sso    bool
}

/* 公共配置 */
func Env() *EnvType {
	c := &EnvType{}
	c.Mode = "dev"                             // 开发环境: dev
	c.ServerHost = "127.0.0.1"                 // 服务器地址
	c.ServerPort = "9030"                      // 服务器端口
	c.Key = "e4b99adec618e653400966be536c45f8" // 加密密钥
	c.Password = "123456"                      // 123456
	// 资源
	c.Img_url = "https://go.webmis.vip/"
	// Token
	c.Admin_token_prefix = "webmisAdmin"       // 前缀-Admin
	c.Admin_token_time = 2 * 3600              // 有效时长(2小时)
	c.Admin_token_auto = true                  // 自动续期
	c.Admin_token_sso = false                  // 单点登录
	c.Api_token_prefix = "webmisApi"           // 前缀-Api
	c.Api_token_time = 7 * 24 * 3600           // 有效时长(7天)
	c.Api_token_auto = true                    // 自动续期
	c.Api_token_sso = true                     // 单点登录
	c.Supplier_token_prefix = "webmisSupplier" // 前缀-Supplier
	c.Supplier_token_time = 7 * 24 * 3600      // 有效时长(7天)
	c.Supplier_token_auto = true               // 自动续期
	c.Supplier_token_sso = true                // 单点登录
	return c
}
