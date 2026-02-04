package core

import (
	"context"
	"time"
	"webmis/app/config"

	"github.com/redis/go-redis/v9"
)

var RedisPool *redis.Client // 连接池
var ctx = context.Background()

/* 控制器 */
type Redis struct {
	Base
	name string
	Conn *redis.Client // 连接
}

/* 初始化 */
func (r *Redis) New(name string) *Redis {
	// 默认值
	r.name = "Pool"
	// 连接池
	if RedisPool == nil {
		// 配置
		cfg := (&config.Redis{}).Config(name)
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
		// 连接池
		RedisPool = redis.NewClient(opt)
		// 校验
		if err := RedisPool.Ping(ctx).Err(); err != nil {
			r.Print("[ "+r.name+" ] Pool:", err.Error())
		}
	}
	// 连接
	if r.Conn == nil {
		r.Conn = RedisPool
	}
	return r
}

/* 添加 */
func (r *Redis) Set(key, value string) bool {
	if r.Conn == nil {
		return false
	}
	return r.Conn.Set(ctx, key, value, 0).Err() == nil
}

/* 自增 */
func (r *Redis) Incr(key string) int64 {
	if r.Conn == nil {
		return 0
	}
	val, _ := r.Conn.Incr(ctx, key).Result()
	return val
}

/* 自减 */
func (r *Redis) Decr(key string) int64 {
	if r.Conn == nil {
		return 0
	}
	val, _ := r.Conn.Decr(ctx, key).Result()
	return val
}

/* 获取 */
func (r *Redis) Get(key string) string {
	if r.Conn == nil {
		return ""
	}
	val, _ := r.Conn.Get(ctx, key).Result()
	return val
}

/* 删除 */
func (r *Redis) Del(key string) bool {
	if r.Conn == nil {
		return false
	}
	return r.Conn.Del(ctx, key).Err() == nil
}

/* 是否存在 */
func (r *Redis) Exist(key string) bool {
	if r.Conn == nil {
		return false
	}
	val, _ := r.Conn.Exists(ctx, key).Result()
	return val > 0
}

/* 设置过期时间(秒) */
func (r *Redis) Expire(key string, expire int64) bool {
	if r.Conn == nil {
		return false
	}
	return r.Conn.Expire(ctx, key, time.Duration(expire)*time.Second).Err() == nil
}

/* 获取过期时间(秒) */
func (r *Redis) TTL(key string) int64 {
	if r.Conn == nil {
		return 0
	}
	val, _ := r.Conn.TTL(ctx, key).Result()
	return int64(val / time.Second)
}

/* 获取长度 */
func (r *Redis) Len(key string) int64 {
	if r.Conn == nil {
		return 0
	}
	val, _ := r.Conn.LLen(ctx, key).Result()
	return val
}

/* 哈希(Hash)-添加 */
func (r *Redis) HSet(key, field, value string) bool {
	if r.Conn == nil {
		return false
	}
	return r.Conn.HSet(ctx, key, field, value).Err() == nil
}

/* 哈希(Hash)-删除 */
func (r *Redis) HDel(key, field string) bool {
	if r.Conn == nil {
		return false
	}
	return r.Conn.HDel(ctx, key, field).Err() == nil
}

/* 哈希(Hash)-获取 */
func (r *Redis) HGet(key, field string) string {
	if r.Conn == nil {
		return ""
	}
	val, _ := r.Conn.HGet(ctx, key, field).Result()
	return val
}

/* 哈希(Hash)-获取全部  */
func (r *Redis) HGetAll(key string) map[string]string {
	if r.Conn == nil {
		return map[string]string{}
	}
	val, _ := r.Conn.HGetAll(ctx, key).Result()
	return val
}

/* 哈希(Hash)-获取全部字段 */
func (r *Redis) HKeys(key string) []string {
	if r.Conn == nil {
		return []string{}
	}
	val, _ := r.Conn.HKeys(ctx, key).Result()
	return val
}

/* 哈希(Hash)-获取全部值 */
func (r *Redis) HVals(key string) []string {
	if r.Conn == nil {
		return []string{}
	}
	val, _ := r.Conn.HVals(ctx, key).Result()
	return val
}

/* 哈希(Hash)-是否存在 */
func (r *Redis) HExist(key, field string) bool {
	if r.Conn == nil {
		return false
	}
	val, _ := r.Conn.HExists(ctx, key, field).Result()
	return val
}

/* 哈希(Hash)-获取长度 */
func (r *Redis) HLen(key string) int64 {
	if r.Conn == nil {
		return 0
	}
	val, _ := r.Conn.HLen(ctx, key).Result()
	return val
}

/* 列表(List)-添加 */
func (r *Redis) LPush(key, value string) bool {
	if r.Conn == nil {
		return false
	}
	return r.Conn.LPush(ctx, key, value).Err() == nil
}
func (r *Redis) RPush(key, value string) bool {
	if r.Conn == nil {
		return false
	}
	return r.Conn.RPush(ctx, key, value).Err() == nil
}

/* 列表(List)-获取 */
func (r *Redis) LPop(key string) string {
	if r.Conn == nil {
		return ""
	}
	val, _ := r.Conn.LPop(ctx, key).Result()
	return val
}
func (r *Redis) RPop(key string) string {
	if r.Conn == nil {
		return ""
	}
	val, _ := r.Conn.RPop(ctx, key).Result()
	return val
}
func (r *Redis) BLPop(key string) string {
	if r.Conn == nil {
		return ""
	}
	val, _ := r.Conn.BLPop(ctx, 0, key).Result()
	return val[1]
}
func (r *Redis) BRPop(key string) string {
	if r.Conn == nil {
		return ""
	}
	val, _ := r.Conn.BRPop(ctx, 0, key).Result()
	return val[1]
}
