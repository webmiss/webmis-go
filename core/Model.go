package core

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"
	"webmis/app/config"
	"webmis/app/util"

	_ "github.com/go-sql-driver/mysql"
)

var Pool *MySQLConnectionPool // 连接池

/* 控制器 */
type Model struct {
	Base
	MySQLConnectionPool
	Conn    *sql.Conn     // 连接
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
func (m *Model) DBConn(name string) *sql.Conn {
	// 默认值
	m.name = "Model"
	m.columns = "*"
	// 配置
	cfg := (&config.Db{}).Config(name)
	// 连接池
	if Pool == nil {
		var err error
		Pool, err = (&MySQLConnectionPool{}).Pool(cfg)
		if err != nil {
			m.Print("[ "+m.name+" ] Pool:", err.Error())
			return nil
		}
		return nil
	}
	// 连接
	conn, err := Pool.GetConnection(3 * time.Second)
	if err != nil {
		m.Print("[ "+m.name+" ] Conn:", err.Error())
		return nil
	}
	m.Conn = conn
	return m.Conn
}

/* 查询 */
func (m *Model) Query(conn *sql.Conn, sql string, args ...interface{}) (*sql.Rows, error) {
	rows, err := m.Conn.QueryContext(context.Background(), sql, args...)
	if err != nil {
		m.Print("[ "+m.name+" ] Query:", err.Error())
		return nil, err
	}
	return rows, nil
}

/* 执行SQL */
func (m *Model) Exec(conn *sql.Conn, sql string, args ...interface{}) sql.Result {
	if conn == nil {
		return nil
	}
	rs, err := conn.ExecContext(context.Background(), sql, args...)
	if err != nil {
		m.Print("[ "+m.name+" ] Exec:", err.Error())
		return nil
	} else {
		nums, _ := rs.RowsAffected()
		m.nums = int(nums)
		return rs
	}
}

/* 关闭 */
func (m *Model) Close() {
	if Pool != nil && m.Conn != nil {
		err := Pool.ReleaseConnection(m.Conn)
		if err != nil {
			m.Print("[ "+m.name+" ] Close:", err.Error())
		}
	}
}

/* 获取-SQL */
func (m *Model) GetSQL() string {
	return m.sql
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
	m.where = " WHERE " + where
	m.args = append(m.args, args...)
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
func (m *Model) Find(sql string, args ...interface{}) []map[string]interface{} {
	// SQL
	if sql == "" {
		sql, args = m.SelectSQL()
		if sql == "" {
			return nil
		}
	}
	// 连接
	if m.Conn == nil {
		return nil
	}
	// 执行
	rows, err := m.Query(m.Conn, sql, args...)
	if err != nil {
		m.Print("[ "+m.name+" ] Find:", err.Error())
		return nil
	}
	return m.FindDataAll(rows)
}

/* 查询-单条 */
func (m *Model) FindFirst(sql string, args ...interface{}) map[string]interface{} {
	// SQL
	if sql == "" {
		m.Limit(0, 1)
		sql, args = m.SelectSQL()
		if sql == "" {
			return nil
		}
	}
	// 连接
	if m.Conn == nil {
		return nil
	}
	// 执行
	rows, err := m.Conn.QueryContext(context.Background(), sql, args...)
	if err != nil {
		m.Print("[ "+m.name+" ] Find:", err.Error())
		return nil
	}
	res := m.FindDataAll(rows)
	if len(res) == 0 {
		return nil
	}
	return res[0]
}

/* 查询-结果 */
func (m *Model) FindDataAll(rs *sql.Rows) []map[string]interface{} {
	res := []map[string]interface{}{}
	// 字段
	columns, _ := rs.Columns()
	key := make([]interface{}, len(columns))
	val := make([]interface{}, len(columns))
	for n := range key {
		key[n] = &val[n]
	}
	// 结果
	for rs.Next() {
		rs.Scan(key...)
		item := map[string]interface{}{}
		for i, v := range val {
			item[columns[i]] = (&util.Type{}).Strval(v)
		}
		res = append(res, item)
	}
	rs.Close()
	m.Close()
	return res
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
func (m *Model) Insert(sql string, args ...interface{}) int {
	if sql == "" {
		sql, args = m.InsertSQL()
		if sql == "" {
			return -1
		}
	}
	rs := m.Exec(m.Conn, sql, args...)
	if rs == nil {
		return -1
	}
	id, _ := rs.LastInsertId()
	m.id = int(id)
	m.Close()
	return m.id
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
	m.sql = "UPDATE " + m.table + " SET " + m.data + m.where
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
func (m *Model) Update(sql string, args ...interface{}) bool {
	if sql == "" {
		sql, args = m.UpdateSQL()
		if sql == "" {
			return false
		}
	}
	rs := m.Exec(m.Conn, sql, args...)
	if rs == nil {
		return false
	}
	m.Close()
	return true
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
	m.sql = "DELETE FROM " + m.table + m.where
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
func (m *Model) Delete(sql string, args ...interface{}) bool {
	if sql == "" {
		sql, args = m.DeleteSQL()
		if sql == "" {
			return false
		}
	}
	rs := m.Exec(m.Conn, sql, args...)
	if rs == nil {
		return false
	}
	m.Close()
	return true
}
