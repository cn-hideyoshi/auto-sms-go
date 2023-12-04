package utils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisUtils struct {
	Host   string
	Port   uint
	Pass   string
	Prefix string
	DB     uint
	Client *redis.Client
}

func (r *RedisUtils) InitClient() {
	r.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Pass,
		DB:       int(r.DB),
	})
}
func (r RedisUtils) buildKey(key string) string {
	return fmt.Sprintf("%s:%s", r.Prefix, key)
}

func (r RedisUtils) Set(key string, value interface{}, expireSec int) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	err := r.Client.Set(ctx, r.buildKey(key), value, time.Duration(expireSec)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r RedisUtils) Get(key string) (string, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	value := r.Client.Get(ctx, r.buildKey(key))
	return value.Result()
}
