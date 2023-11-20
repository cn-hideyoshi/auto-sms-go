package main

import (
	_ "blog.hideyoshi.top/company/internal/cache"
	_ "blog.hideyoshi.top/company/internal/db/dao"
	"blog.hideyoshi.top/company/server"
)

func main() {
	server.Start()
}
