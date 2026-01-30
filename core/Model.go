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
	keys    string        // 添加-键
	values  string        // 添加-值
	data    string        // 更新-数据
	id      int           // 自增ID
	nums    int           // 影响行数
}

/* 获取连接 */
func (m *Model) DBConn(name string) *sql.DB {
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
	return m.Conn
}

/* 执行SQL */
func (m *Model) Exec(conn *sql.DB, sql string, args []interface{}) sql.Result {
	rows, err := conn.Exec(sql, args...)
	if err != nil {
		m.Print("[ "+m.name+" ] Exec:", err.Error())
		return nil
	} else {
		id, _ := rows.LastInsertId()
		nums, _ := rows.RowsAffected()
		m.id = int(id)
		m.nums = int(nums)
		return rows
	}
}

/* 获取-SQL */
func (m *Model) GetSQL() (string, []interface{}) {
	return m.sql, m.args
}

/* 获取-自增ID */
func (m *Model) GetID() int {
	return m.id
}

/* 获取-影响行数 */
func (m *Model) GetNums() int {
	return m.nums
}

/* 表名 */
func (m *Model) Table(name string) {
	m.table = name
}

/* 分区 */
func (m *Model) Partition(partition ...string) {
	m.table = m.table + " PARTITION('" + strings.Join(partition, ",") + "')"
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
	m.args = args
}

/* 分组 */
func (m *Model) Group(group ...string) {
	m.group = " GROUP BY " + strings.Join(group, ",")
}

/* 筛选 */
func (m *Model) Having(having string) {
	m.having = " HAVING " + having
}

/* 排序 */
func (m *Model) Order(order ...string) {
	m.order = " ORDER BY " + strings.Join(order, ",")
}

/* 限制 */
func (m *Model) Limit(start int, limit int) {
	m.limit = " LIMIT " + strconv.FormatInt(int64(start), 10) + "," + strconv.FormatInt(int64(limit), 10)
}

/* 分页 */
func (m *Model) Page(page int, limit int) {
	m.limit = " LIMIT " + strconv.FormatInt(int64((page-1)*limit), 10) + "," + strconv.FormatInt(int64(limit), 10)
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
func (m *Model) Find(sql string, args []interface{}) {
	if sql == "" {
		sql, args = m.SelectSQL()
	}
	m.Print(sql, args)
}

/* 查询-单条 */
func (m *Model) FindFirst(sql string, args []interface{}) {
	if sql == "" {
		m.Limit(0, 1)
		sql, args = m.SelectSQL()
	}
	m.Print(sql, args)
}

/* 添加-单条 */
func (m *Model) Values(data map[string]interface{}) {
	m.args = []interface{}{}
	var keys, vals string
	for k, v := range data {
		keys += k + ","
		vals += "?,"
		m.args = append(m.args, v)
	}
	m.keys = strings.TrimRight(keys, ",")
	m.values = strings.TrimRight(vals, ",")
}

/* 添加-多条 */
func (m *Model) ValuesAll(data []map[string]interface{}) {
	m.args = []interface{}{}
	var keys, vals, tmp string
	for _, v := range data {
		tmp = ""
		for k, val := range v {
			keys += k + ","
			tmp += "?,"
			m.args = append(m.args, val)
		}
		vals += "(" + strings.TrimRight(tmp, ",") + "),"
	}
	m.keys = strings.TrimRight(keys, ",")
	m.values = strings.TrimRight(vals, ",")
}

/* 添加-SQL */
func (m *Model) InsertSQL() (string, []interface{}) {
	// 验证
	if m.table == "" {
		m.Print("[ "+m.name+" ]", "Insert: 表不能为空!")
		return "", nil
	}
	if m.keys == "" || m.values == "" {
		m.Print("[ "+m.name+" ]", "Insert: 字段或值不能为空!")
		return "", nil
	}
	// SQL
	m.sql = "INSERT INTO " + m.table + " (" + m.keys + ") VALUES(" + m.values + ")"
	m.table = ""
	m.keys = ""
	m.values = ""
	// 参数
	args := m.args
	m.args = make([]interface{}, 0, 10)
	// 结果
	return m.sql, args
}

/* 添加-执行 */
func (m *Model) Insert(sql string, args []interface{}) {
	if sql == "" {
		sql, args = m.InsertSQL()
	}
	m.Print(sql, args)
}

/* 更新-数据 */
func (m *Model) Set(data map[string]interface{}) {
	m.args = []interface{}{}
	var vals string
	for k, v := range data {
		vals += k + "=?,"
		m.args = append(m.args, v)
	}
	m.data = strings.TrimRight(vals, ",")
}

/* 更新-SQL */
func (m *Model) UpdateSQL() (string, []interface{}) {
	// 验证
	if m.table == "" {
		m.Print("[ "+m.name+" ]", "Update: 表不能为空!")
		return "", nil
	}
	if m.data == "" {
		m.Print("[ "+m.name+" ]", "Update: 数据不能为空!")
		return "", nil
	}
	if m.where == "" {
		m.Print("[ "+m.name+" ]", "Update: 条件不能为空!")
		return "", nil
	}
	// SQL
	m.sql = "UPDATE " + m.table + " SET " + m.data + " WHERE " + m.where
	// 重置
	m.table = ""
	m.data = ""
	m.where = ""
	// 参数
	args := m.args
	m.args = make([]interface{}, 0, 10)
	// 结果
	return m.sql, args
}

/* 更新-执行 */
func (m *Model) Update(sql string, args []interface{}) {
	if sql == "" {
		sql, args = m.UpdateSQL()
	}
	m.Print(sql, args)
}

/* 删除-SQL */
func (m *Model) DeleteSQL() (string, []interface{}) {
	// 验证
	if m.table == "" {
		m.Print("[ "+m.name+" ]", "Delete: 表不能为空!")
		return "", nil
	}
	if m.where == "" {
		m.Print("[ "+m.name+" ]", "Delete: 条件不能为空!")
		return "", nil
	}
	// SQL
	m.sql = "DELETE FROM " + m.table + " WHERE " + m.where
	// 重置
	m.table = ""
	m.where = ""
	// 参数
	args := m.args
	m.args = make([]interface{}, 0, 10)
	// 结果
	return m.sql, args
}

/* 删除-执行 */
func (m *Model) Delete(sql string, args []interface{}) {
	if sql == "" {
		sql, args = m.DeleteSQL()
	}
	m.Print(sql, args)
}
