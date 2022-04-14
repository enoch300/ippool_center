package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	RDB *redis.Client
)

// Connect 初始化连接
func Connect() (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = RDB.Ping(c).Result()
	return err
}
