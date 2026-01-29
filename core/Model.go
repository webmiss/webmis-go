package core

import (
	"database/sql"
	"strconv"
	"strings"
	"webmis/app/config"

	_ "github.com/go-sql-driver/mysql"
)

/* 控制器 */
type Model struct {
	Base
	Conn    *sql.DB       // 连接
	name    string        // 名称
	table   string        // 数据表
	columns string        // 字段
	where   string        // 条件
	group   string        // 分组
	having  string        // 筛选
	order   string        // 排序
	limit   string        // 限制
	args    []interface{} // 参数
	sql     string        // SQL语句
}

/* 获取连接 */
func (m *Model) DBConn(name string) bool {
	// 默认值
	m.name = "Model"
	m.columns = "*"
	// 配置
	cfg := (&config.Db{}).Config(name)
	if m.Conn == nil {
		dsn := cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Database + "?charset=" + cfg.Charset + "&parseTime=True&loc=" + cfg.Loc
		m.Print(dsn)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			m.Print("[ "+m.name+" ] Conn:", err.Error())
		}
		m.Conn = db
	}
	if m.Conn != nil {
		return true
	}
	return false
}

/* 表名 */
func (m *Model) Table(name string) {
	m.table = name
	m.Table("user")
}

/* 分区 */
func (m *Model) Partition(partition ...string) {
	m.table = m.table + "PARTITION('" + strings.Join(partition, ",") + "')"
}

/* 关联-INNER */
func (m *Model) InnerJoin(table string, on string) {
	m.table = m.table + " INNER JOIN " + table + " ON " + on
}

/* 关联-LEFT */
func (m *Model) LeftJoin(table string, on string) {
	m.table = m.table + " LEFT JOIN " + table + " ON " + on
}

/* 关联-RIGHT */
func (m *Model) RightJoin(table string, on string) {
	m.table = m.table + " RIGHT JOIN " + table + " ON " + on
}

/* 关联-FULL */
func (m *Model) FullJoin(table string, on string) {
	m.table = m.table + " FULL JOIN " + table + " ON " + on
}

/* 字段 */
func (m *Model) Columns(columns ...string) {
	m.columns = strings.Join(columns, ",")
}

/* 条件 */
func (m *Model) Where(where string, args ...interface{}) {
	m.where = where
	m.args = append(m.args, args...)
}

/* 分组 */
func (m *Model) Group(group ...string) {
	m.group = "GROUP BY " + strings.Join(group, ",")
}

/* 筛选 */
func (m *Model) Having(having string) {
	m.having = "HAVING " + having
}

/* 排序 */
func (m *Model) Order(order ...string) {
	m.order = "ORDER BY " + strings.Join(order, ",")
}

/* 限制 */
func (m *Model) Limit(start int, limit int) {
	m.limit = "LIMIT " + strconv.FormatInt(int64(start), 10) + "," + strconv.FormatInt(int64(limit), 10)
}

/* 分页 */
func (m *Model) Page(page int, limit int) {
	m.limit = "LIMIT " + strconv.FormatInt(int64((page-1)*limit), 10) + "," + strconv.FormatInt(int64(limit), 10)
}

/* 查询-SQL */
func (m *Model) SelectSQL() (string, []interface{}) {
	// 验证
	if m.table == "" {
		m.Print("[ "+m.name+" ]", "Select: 表不能为空!")
		return "", nil
	}
	if m.columns == "" {
		m.Print("[ "+m.name+" ]", "Select: 表不能为空!")
		return "", nil
	}
	// SQL
	m.sql = "SELECT " + m.columns + " FROM " + m.table
	m.table = ""
	m.columns = "*"
	if m.where != "" {
		m.sql += m.where
		m.where = ""
	}
	if m.group != "" {
		m.sql += m.group
		m.group = ""
	}
	if m.having != "" {
		m.sql += m.having
		m.having = ""
	}
	if m.order != "" {
		m.sql += m.order
		m.order = ""
	}
	if m.limit != "" {
		m.sql += m.limit
		m.limit = ""
	}
	// 参数
	args := m.args
	m.args = make([]interface{}, 0, 10)
	// 结果
	return m.sql, args
}

/* 查询-多条 */
func (m *Model) Find(param []interface{}) {
	sql, args := m.SelectSQL()
	if param[0] != nil {
		sql = param[0].(string)
	}
	if param[1] != nil {
		args = param[0].([]interface{})
	}
	m.Print(sql, args)
}
