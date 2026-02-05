package config

/* 缓存数据库 */
type Redis struct {
	Host     string // 主机
	Port     string // 端口
	Password string // 密码
	Db       int    // 硬盘
	MaxTotal int    // 最大连接数
	MaxIdle  int    // 最大空闲连接数
	MinIdle  int    // 最小空闲连接数
	MaxWait  int    // 最大等待时间
}

/* 配置 */
func (c *Redis) Config(name string) *Redis {
	switch name {
	case "default":
		c.Host = "127.0.0.1"                            // 主机
		c.Port = "6379"                                 // 端口
		c.Password = "e4b99adec618e653400966be536c45f8" // 密码
		c.Db = 0                                        // 硬盘
		c.MaxTotal = 30                                 // 最大连接数
		c.MinIdle = 10                                  // 最大空闲连接数
		c.MaxIdle = 15                                  // 最大空闲连接数
		c.MaxWait = 3                                   // 最大等待时间
	case "other":
		c.Host = "127.0.0.1"                            // 主机
		c.Port = "6379"                                 // 端口
		c.Password = "e4b99adec618e653400966be536c45f8" // 密码
		c.Db = 0                                        // 硬盘
		c.MaxTotal = 30                                 // 最大连接数
		c.MinIdle = 10                                  // 最大空闲连接数
		c.MaxIdle = 15                                  // 最大空闲连接数
		c.MaxWait = 3                                   // 最大等待时间
	}
	return c
}
