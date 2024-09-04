package manage

import (
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

type RedisManage struct {
	Client *redis.Client
}

var redisManageOnce sync.Once
var redisManageCache *RedisManage

func init() {
	redisManageCache = &RedisManage{}
}
func GetRedisManage() *RedisManage {
	redisManageOnce.Do(func() {
		redisManageCache.Client = redis.NewClient(&redis.Options{
			Addr:         GetConfigManage().GetString("redis.host") + ":" + GetConfigManage().GetString("redis.port"),
			Password:     GetConfigManage().GetString("redis.password"),
			DB:           GetConfigManage().GetInt("redis.db"),
			PoolSize:     GetConfigManage().GetInt("redis.pool_size"),
			MinIdleConns: GetConfigManage().GetInt("redis.min_free_conn"),
			DialTimeout:  1 * time.Second,
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 1 * time.Second,
		})
	})
	return redisManageCache
}
