package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Database struct {
	Driver         string
	DataSourceName string
	Db             *sqlx.DB
}

func (db *Database) InitDb() error {
	var err error
	db.Db, err = sqlx.Connect(db.Driver, db.DataSourceName)
	if err != nil {
		log.Fatalln(err)
	}
	db.Db.SetMaxIdleConns(10)
	db.Db.SetMaxOpenConns(100)
	db.Db.SetConnMaxLifetime(time.Second * 30)
	return nil
}
