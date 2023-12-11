package dao

import (
	"blog.hideyoshi.top/common/pkg/db/dao"
	"blog.hideyoshi.top/msg/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

var _db *sqlx.DB

func init() {
	DbConfig := config.C.Db
	DataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", DbConfig.Username, DbConfig.Password, DbConfig.Host, DbConfig.Port, DbConfig.DbName)
	database := dao.Database{
		Driver:         DbConfig.Driver,
		DataSourceName: DataSourceName,
	}
	err := database.InitDb()
	if err != nil {
		log.Println("init db fail", err)
	}
	_db = database.Db
}
