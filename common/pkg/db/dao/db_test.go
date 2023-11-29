package dao

import (
	"fmt"
	"log"
	"testing"
)

type TestTb struct {
	Id int64 `db:"id" json:"id"`
}

func TestDatabase_InitDb(t *testing.T) {
	database := Database{
		Driver:         "mysql",
		DataSourceName: "hideyoshi:123456@tcp(192.168.1.5:3306)/welfare-go",
	}
	err := database.InitDb()
	if err != nil {
		log.Fatalln(err)
	}

	queryx, err := database.Db.Queryx("select * from as_test")
	if err != nil {
		log.Fatalln(err)
	}
	for queryx.Next() {
		testTb := &TestTb{}
		err := queryx.StructScan(testTb)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", testTb)
	}
}
