package core

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"webmis/app/config"

	_ "github.com/go-sql-driver/mysql"
)

var pool_default chan *sql.Conn  // 连接池: default
var pool_other chan *sql.Conn    // 连接池: other
var pool_name string = "MariaDB" // 名称
var pool_db string = "default"   // 数据库
var pool_initSize int            // 初始连接数
var pool_maxSize int             // 最大连接数
var pool_maxWait time.Duration   // 超时时间
var pool_dsn string              // MySQL连接DSN

// MySQL 连接池
type MySQLConnectionPool struct {
	Base
}

/* 数据源 */
func (p *MySQLConnectionPool) InitPool(name string) {
	// 参数
	pool_db = name
	// 配置
	cfg := (&config.Db{}).Config(name)
	pool_initSize = cfg.PoolInitSize
	pool_maxSize = cfg.PoolMaxSize
	pool_maxWait = time.Duration(cfg.PoolMaxWait) * time.Millisecond
	pool_dsn = cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Database + "?charset=" + cfg.Charset + "&parseTime=True&loc=" + cfg.Loc
	// 初始化连接池
	if name == "default" && pool_default != nil {
		return
	}
	if name == "other" && pool_other != nil {
		return
	}
	// 创建连接池
	if name == "default" {
		pool_default = make(chan *sql.Conn, cfg.PoolMaxSize)
	} else if name == "other" {
		pool_other = make(chan *sql.Conn, cfg.PoolMaxSize)
	}
	// 初始化连接数
	for i := 0; i < pool_initSize; i++ {
		conn, err := p.CreateConnection()
		if err != nil {
			p.Print("[ " + pool_name + " ] MariaDB Pool: " + err.Error())
			return
		}
		if name == "default" {
			pool_default <- conn
		} else if name == "other" {
			pool_other <- conn
		}
	}
	p.Print("[ "+pool_name+" ] MariaDB Pool:", pool_db, p.GetIdleCount())
}

/* 创建连接 */
func (p *MySQLConnectionPool) CreateConnection() (*sql.Conn, error) {
	db, err := sql.Open("mysql", pool_dsn)
	if err != nil {
		return nil, err
	}
	// 设置属性
	db.SetMaxIdleConns(pool_maxSize)
	db.SetMaxOpenConns(pool_maxSize)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	// 获取连接
	conn, err := db.Conn(context.Background())
	if err != nil {
		_ = db.Close()
		p.Print("[ " + pool_name + " ] CreateConnection: " + err.Error())
		return nil, err
	}
	return conn, nil
}

/* 默认连接池 */
func (p *MySQLConnectionPool) GetIdleConnections() chan *sql.Conn {
	if pool_db == "default" {
		return pool_default
	} else if pool_db == "other" {
		return pool_other
	}
	return nil
}

/* 获取连接 */
func (p *MySQLConnectionPool) GetConnection() (*sql.Conn, error) {
	idle := p.GetIdleConnections()
	if idle == nil {
		return nil, errors.New("[ " + pool_name + " ]" + " 无效连接池: " + pool_db)
	}
	// 超时获取连接
	ctx, cancel := context.WithTimeout(context.Background(), pool_maxWait)
	defer cancel()
	select {
	case conn := <-idle:
		// 从空闲通道获取连接
		if !p.ValidateConn(conn) {
			_ = conn.Close()
			return p.CreateConnection()
		}
		return conn, nil
	case <-ctx.Done():
		// 验证
		if p.GetIdleCount() < pool_maxSize {
			conn, err := p.CreateConnection()
			if err != nil {
				return nil, errors.New("[ " + pool_name + " ]" + "GetConnection: " + err.Error())
			}
			return conn, nil
		}
		return nil, errors.New("[ " + pool_name + " ]" + "Connection pool is full, timeout while acquiring idle connection")
	}
}

/* 归还连接 */
func (p *MySQLConnectionPool) ReleaseConnection(conn *sql.Conn) bool {
	if conn == nil {
		return false
	}
	// 连接池
	idleConnections := p.GetIdleConnections()
	if idleConnections == nil {
		return false
	}
	// 校验有效性
	if !p.ValidateConn(conn) {
		_ = conn.Close()
		return false
	}
	// 归还连接
	select {
	case idleConnections <- conn:
		return true
	default:
		_ = conn.Close()
		p.Print("[ " + pool_name + " ] ReleaseConnection: 连接池已满")
		return false
	}
}

/* 验证连接 */
func (p *MySQLConnectionPool) ValidateConn(conn *sql.Conn) bool {
	if conn == nil {
		return false
	}
	err := conn.PingContext(context.Background())
	if err != nil {
		p.Print("[ " + pool_name + " ] ValidateConn: " + err.Error())
		return false
	}
	return true
}

/* 获取空闲连接数 */
func (p *MySQLConnectionPool) GetIdleCount() int {
	// 连接池
	idleConnections := p.GetIdleConnections()
	if idleConnections == nil {
		return 0
	}
	return len(idleConnections)
}

/* 销毁连接池 */
func (p *MySQLConnectionPool) Destroy() {
	// 连接池: default
	if pool_default != nil {
		close(pool_default)
		for conn := range pool_default {
			if conn != nil {
				_ = conn.Close()
			}
		}
		pool_default = nil
	}
	// 连接池: other
	if pool_other != nil {
		close(pool_other)
		for conn := range pool_other {
			if conn != nil {
				_ = conn.Close()
			}
		}
		pool_other = nil
	}
}
