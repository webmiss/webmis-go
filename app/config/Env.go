package config

type EnvType struct {
	Mode       string
	ServerHost string
	ServerPort string
}

/* 公共配置 */
func Env() *EnvType {
	cfg := &EnvType{}
	cfg.Mode = "dev"             // 开发环境: dev
	cfg.ServerHost = "127.0.0.1" // 服务器地址
	cfg.ServerPort = "9030"      // 服务器端口
	return cfg
}
