package cache

import (
	"blog.hideyoshi.top/common/utils"
	"blog.hideyoshi.top/company/config"
	"fmt"
)

var Cache *utils.RedisUtils

func init() {
	Cache = &utils.RedisUtils{
		Host:   config.C.Redis.Host,
		Port:   config.C.Redis.Port,
		Prefix: config.C.Redis.Prefix,
		Pass:   config.C.Redis.Pass,
		DB:     config.C.Redis.DB,
	}
	Cache.InitClient()
}

var keyPre = "company:"

func buildKey(key string) string {
	return fmt.Sprintf("%s%s", keyPre, key)
}
