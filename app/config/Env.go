package config

type EnvType struct {
	Mode       string
	ServerHost string
	ServerPort string
	Key        string
	Password   string
}

/* 公共配置 */
func Env() *EnvType {
	cfg := &EnvType{}
	cfg.Mode = "dev"                             // 开发环境: dev
	cfg.ServerHost = "127.0.0.1"                 // 服务器地址
	cfg.ServerPort = "9030"                      // 服务器端口
	cfg.Key = "e4b99adec618e653400966be536c45f8" // 加密密钥
	cfg.Password = "123456"                      // 123456
	return cfg
}
