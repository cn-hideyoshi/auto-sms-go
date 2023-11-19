package model

import "time"

type User struct {
	Id           int64     `db:"id" json:"id"`
	UserName     string    `db:"user_name" json:"userName"`
	UserPassword string    `db:"user_password" json:"userPassword"`
	CreateTime   time.Time `db:"create_time" json:"create_time"`
	UpdateTime   time.Time `db:"update_time" json:"update_time"`
}
