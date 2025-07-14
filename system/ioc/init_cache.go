package ioc

import "github.com/redis/go-redis/v9"

func InitRedis() redis.Cmdable {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	return rdb
}
