package utils

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisUtils struct {
	Host   string
	Port   uint
	Pass   string
	DB     uint
	Client *redis.Client
}

func (r *RedisUtils) InitClient() {
	r.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Pass,    // 没有密码
		DB:       int(r.DB), // 使用默认数据库
	})
}
