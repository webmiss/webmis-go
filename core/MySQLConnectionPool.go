package core

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"webmis/app/config"

	_ "github.com/go-sql-driver/mysql"
)

// MySQL 连接池
type MySQLConnectionPool struct {
	name      string         // 名称
	idleConns chan *sql.Conn // 存储空闲连接的通道（并发安全）
	maxSize   int            // 连接池最大连接数
	initSize  int            // 初始连接数
	dsn       string         // MySQL连接DSN
	closed    bool           // 连接池是否已关闭
	connCount int            // 已创建的总连接数（含使用中+空闲）
	countLock chan struct{}  // 用于保护connCount的锁（基于通道的轻量锁）
}

/* 初始化连接池 */
func (p *MySQLConnectionPool) Pool(cfg *config.Db) (*MySQLConnectionPool, error) {
	// 参数
	p.name = "Pool"
	p.closed = false
	p.countLock = make(chan struct{}, 1)
	p.initSize = cfg.PoolInitSize
	p.maxSize = cfg.PoolMaxSize
	if p.initSize > p.maxSize {
		p.maxSize = p.initSize
	}
	// 配置
	p.idleConns = make(chan *sql.Conn, cfg.PoolMaxSize)
	p.dsn = cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Database + "?charset=" + cfg.Charset + "&parseTime=True&loc=" + cfg.Loc
	// 创建初始连接
	for i := 0; i < p.initSize; i++ {
		conn, err := p.createConn()
		if err != nil {
			p.Destroy()
			return nil, errors.New("[ " + p.name + " ]" + "创建初始连接失败: " + err.Error())
		}
		p.idleConns <- conn
		p.incrConnCount()
	}
	return p, nil
}

/* 创建连接 */
func (p *MySQLConnectionPool) createConn() (*sql.Conn, error) {
	if p.closed {
		return nil, errors.New("[ " + p.name + " ]" + "连接池已关闭，无法创建新连接")
	}
	db, err := sql.Open("mysql", p.dsn)
	if err != nil {
		return nil, errors.New("[ " + p.name + " ]" + "创建连接失败: " + err.Error())
	}
	// 设置属性
	db.SetMaxIdleConns(p.maxSize)
	db.SetMaxOpenConns(p.maxSize)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	// 获取连接
	conn, err := db.Conn(context.Background())
	if err != nil {
		db.Close()
		return nil, errors.New("[ " + p.name + " ]" + "获取连接失败: " + err.Error())
	}
	// 是否有效
	if err := p.validateConn(conn); err != nil {
		conn.Close()
		return nil, errors.New("[ " + p.name + " ]" + "连接无效: " + err.Error())
	}
	return conn, nil
}

/* 获取连接 */
func (p *MySQLConnectionPool) GetConnection(timeout time.Duration) (*sql.Conn, error) {
	if p.closed {
		return nil, errors.New("[ " + p.name + " ]" + "连接池已关闭，无法获取连接")
	}
	// 超时获取连接
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	select {
	case <-ctx.Done():
		// 验证
		if p.getConnCount() < p.maxSize {
			conn, err := p.createConn()
			if err != nil {
				return nil, errors.New("[ " + p.name + " ]" + "创建连接失败: " + err.Error())
			}
			p.incrConnCount()
			return conn, nil
		}
		return nil, errors.New("[ " + p.name + " ]" + "连接池已满，无法获取连接")
	case conn := <-p.idleConns:
		// 从空闲通道获取连接
		if err := p.validateConn(conn); err != nil {
			conn.Close()
			p.decrConnCount()
			return p.createConn()
		}
		return conn, nil
	}
}

/* 归还连接 */
func (p *MySQLConnectionPool) ReleaseConnection(conn *sql.Conn) error {
	if p.closed || conn == nil {
		if conn != nil {
			conn.Close()
		}
		return nil
	}
	// 校验有效性
	if err := p.validateConn(conn); err != nil {
		conn.Close()
		p.decrConnCount()
		return errors.New("[ " + p.name + " ]" + "连接无效: " + err.Error())
	}
	// 归还连接
	select {
	case p.idleConns <- conn:
		return nil
	default:
		conn.Close()
		p.decrConnCount()
		return errors.New("[ " + p.name + " ]" + "连接池已满，无法归还连接")
	}
}

/* 验证连接 */
func (p *MySQLConnectionPool) validateConn(conn *sql.Conn) error {
	return conn.PingContext(context.Background())
}

/* 线程安全-增加连接计数 */
func (p *MySQLConnectionPool) incrConnCount() {
	p.countLock <- struct{}{}
	p.connCount++
	<-p.countLock
}

/* 线程安全-减少连接计数 */
func (p *MySQLConnectionPool) decrConnCount() {
	p.countLock <- struct{}{}
	p.connCount--
	<-p.countLock
}

/* 线程安全-获取空闲连接数 */
func (p *MySQLConnectionPool) getConnCount() int {
	p.countLock <- struct{}{}
	count := p.connCount
	<-p.countLock
	return count
}

/* 获取空闲连接数 */
func (p *MySQLConnectionPool) GetIdleCount() int {
	return len(p.idleConns)
}

/* 销毁连接池 */
func (p *MySQLConnectionPool) Destroy() {
	if p.closed {
		return
	}
	p.closed = true
	// 关闭所有空闲连接
	close(p.idleConns)
	for conn := range p.idleConns {
		if conn != nil {
			conn.Close()
		}
	}
	// 重置计数
	p.connCount = 0
}
