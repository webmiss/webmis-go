package config

type EnvType struct {
	ServerHost string
	ServerPort string
}

/* 公共配置 */
func Env() *EnvType {
	cfg := &EnvType{}
	cfg.ServerHost = "127.0.0.1"
	cfg.ServerPort = "9030"
	return cfg
}
