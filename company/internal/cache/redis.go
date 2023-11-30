package cache

import (
	"blog.hideyoshi.top/common/utils"
	"blog.hideyoshi.top/company/config"
	"context"
	"fmt"
	"time"
)

var Cache *utils.RedisUtils

func init() {
	Cache = &utils.RedisUtils{
		Host: config.C.Redis.Host,
		Port: config.C.Redis.Port,
		Pass: config.C.Redis.Pass,
		DB:   config.C.Redis.DB,
	}
	Cache.InitClient()
}

var keyPre = "company:"

func buildKey(key string) string {
	return fmt.Sprintf("%s%s", keyPre, key)
}
func Set(key string, value interface{}, expireSec int) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	err := Cache.Client.Set(ctx, buildKey(key), value, time.Duration(expireSec)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) (string, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	value := Cache.Client.Get(ctx, buildKey(key))
	return value.Result()
}
