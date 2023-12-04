package cache

import (
	"blog.hideyoshi.top/common/utils"
	"blog.hideyoshi.top/user/config"
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
