package config

/* 数据库 */
type Db struct {
	Host     string // 主机
	Port     string // 端口
	User     string // 用户名
	Password string // 密码
	Database string // 数据库
	Charset  string // 编码
	Loc      string // 时区
}

/* 配置 */
func (c *Db) Config(name string) *Db {
	switch name {
	case "default":
		c.Host = "127.0.0.1"  // 主机
		c.Port = "3306"       // 端口
		c.User = "root"       // 用户名
		c.Password = "123456" // 密码
		c.Database = "webmis" // 数据库
		c.Charset = "utf8mb4" // 编码
		c.Loc = "Local"       // 时区
	case "other":
		c.Host = "127.0.0.1"                            // 主机
		c.Port = "3306"                                 // 端口
		c.User = "root"                                 // 用户名
		c.Password = "e4b99adec618e653400966be536c45f8" // 密码
		c.Database = "webmis"                           // 数据库
		c.Charset = "utf8mb4"                           // 编码
		c.Loc = "Local"                                 // 时区
	}
	return c
}
