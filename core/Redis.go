package core

import (
	"context"
	"time"
	"webmis/app/config"

	"github.com/redis/go-redis/v9"
)

var redis_default *redis.Client // 连接池: default
var redis_other *redis.Client   // 连接池: other
var redis_db string = "default" // 数据库
var ctx = context.Background()

/* 控制器 */
type Redis struct {
	Base
	name string
}

/* 初始化 */
func (r *Redis) New(name string) *Redis {
	// 参数
	r.name = "Redis"
	// 数据库
	if name == "" {
		name = "default"
	}
	redis_db = name
	// 连接池
	if name == "default" && redis_default != nil {
		return r
	}
	if name == "other" && redis_other != nil {
		return r
	}
	// 初始化连接池
	cfg := (&config.Redis{}).Config(name)
	// 配置
	opt := &redis.Options{
		Addr:         cfg.Host + ":" + cfg.Port,
		Password:     cfg.Password,
		DB:           cfg.Db,
		PoolSize:     cfg.MaxTotal,
		MinIdleConns: cfg.MinIdle,
		MaxIdleConns: cfg.MaxIdle,
		PoolTimeout:  time.Duration(cfg.MaxWait) * time.Second,
		DialTimeout:  time.Duration(cfg.MaxWait) * time.Second,
		ReadTimeout:  time.Duration(cfg.MaxWait) * time.Second,
		WriteTimeout: time.Duration(cfg.MaxWait) * time.Second,
	}
	// 创建连接池
	if name == "default" {
		redis_default = redis.NewClient(opt)
	}
	if name == "other" {
		redis_other = redis.NewClient(opt)
	}
	r.Print("[ "+r.name+" ] Redis Pool:", name, cfg.MaxTotal)
	return r
}

/* 获取连接 */
func (r *Redis) RedisConn() *redis.Client {
	if redis_db == "default" {
		return redis_default
	} else if redis_db == "other" {
		return redis_other
	}
	return nil
}

/* 添加 */
func (r *Redis) Set(key, value string) bool {
	conn := r.RedisConn()
	if conn == nil {
		return false
	}
	return conn.Set(ctx, key, value, 0).Err() == nil
}

/* 自增 */
func (r *Redis) Incr(key string) int64 {
	conn := r.RedisConn()
	if conn == nil {
		return 0
	}
	val, _ := conn.Incr(ctx, key).Result()
	return val
}

/* 自减 */
func (r *Redis) Decr(key string) int64 {
	conn := r.RedisConn()
	if conn == nil {
		return 0
	}
	val, _ := conn.Decr(ctx, key).Result()
	return val
}

/* 获取 */
func (r *Redis) Get(key string) string {
	conn := r.RedisConn()
	if conn == nil {
		return ""
	}
	val, _ := conn.Get(ctx, key).Result()
	return val
}

/* 删除 */
func (r *Redis) Del(key string) bool {
	conn := r.RedisConn()
	if conn == nil {
		return false
	}
	return conn.Del(ctx, key).Err() == nil
}

/* 是否存在 */
func (r *Redis) Exist(key string) bool {
	conn := r.RedisConn()
	if conn == nil {
		return false
	}
	val, _ := conn.Exists(ctx, key).Result()
	return val > 0
}

/* 设置过期时间(秒) */
func (r *Redis) Expire(key string, expire int64) bool {
	conn := r.RedisConn()
	if conn == nil {
		return false
	}
	return conn.Expire(ctx, key, time.Duration(expire)*time.Second).Err() == nil
}

/* 获取过期时间(秒) */
func (r *Redis) Ttl(key string) int64 {
	conn := r.RedisConn()
	if conn == nil {
		return 0
	}
	val, _ := conn.TTL(ctx, key).Result()
	return int64(val / time.Second)
}

/* 获取长度 */
func (r *Redis) StrLen(key string) int64 {
	conn := r.RedisConn()
	if conn == nil {
		return 0
	}
	val, _ := conn.StrLen(ctx, key).Result()
	return val
}

/* 哈希(Hash)-添加 */
func (r *Redis) HSet(key, field, value string) bool {
	conn := r.RedisConn()
	if conn == nil {
		return false
	}
	return conn.HSet(ctx, key, field, value).Err() == nil
}

/* 哈希(Hash)-删除 */
func (r *Redis) HDel(key, field string) bool {
	conn := r.RedisConn()
	if conn == nil {
		return false
	}
	return conn.HDel(ctx, key, field).Err() == nil
}

/* 哈希(Hash)-获取 */
func (r *Redis) HGet(key, field string) string {
	conn := r.RedisConn()
	if conn == nil {
		return ""
	}
	val, _ := conn.HGet(ctx, key, field).Result()
	return val
}

/* 哈希(Hash)-获取全部  */
func (r *Redis) HGetAll(key string) map[string]string {
	conn := r.RedisConn()
	if conn == nil {
		return map[string]string{}
	}
	val, _ := conn.HGetAll(ctx, key).Result()
	return val
}

/* 哈希(Hash)-获取全部字段 */
func (r *Redis) HKeys(key string) []string {
	conn := r.RedisConn()
	if conn == nil {
		return []string{}
	}
	val, _ := conn.HKeys(ctx, key).Result()
	return val
}

/* 哈希(Hash)-获取全部值 */
func (r *Redis) HVals(key string) []string {
	conn := r.RedisConn()
	if conn == nil {
		return []string{}
	}
	val, _ := conn.HVals(ctx, key).Result()
	return val
}

/* 哈希(Hash)-是否存在 */
func (r *Redis) HExist(key, field string) bool {
	conn := r.RedisConn()
	if conn == nil {
		return false
	}
	val, _ := conn.HExists(ctx, key, field).Result()
	return val
}

/* 哈希(Hash)-获取长度 */
func (r *Redis) HLen(key string) int64 {
	conn := r.RedisConn()
	if conn == nil {
		return 0
	}
	val, _ := conn.HLen(ctx, key).Result()
	return val
}

/* 列表(List)-添加 */
func (r *Redis) LPush(key, value string) bool {
	conn := r.RedisConn()
	if conn == nil {
		return false
	}
	return conn.LPush(ctx, key, value).Err() == nil
}
func (r *Redis) RPush(key, value string) bool {
	conn := r.RedisConn()
	if conn == nil {
		return false
	}
	return conn.RPush(ctx, key, value).Err() == nil
}

/* 列表(List)-获取 */
func (r *Redis) LPop(key string) string {
	conn := r.RedisConn()
	if conn == nil {
		return ""
	}
	val, _ := conn.LPop(ctx, key).Result()
	return val
}
func (r *Redis) RPop(key string) string {
	conn := r.RedisConn()
	if conn == nil {
		return ""
	}
	val, _ := conn.RPop(ctx, key).Result()
	return val
}
func (r *Redis) BLPop(key string) []string {
	conn := r.RedisConn()
	if conn == nil {
		return []string{}
	}
	val, _ := conn.BLPop(ctx, 0, key).Result()
	return val
}
func (r *Redis) BRPop(key string) []string {
	conn := r.RedisConn()
	if conn == nil {
		return []string{}
	}
	val, _ := conn.BRPop(ctx, 0, key).Result()
	return val
}
