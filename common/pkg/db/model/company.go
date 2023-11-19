package model

import "time"

type Company struct {
	CompanyId       int64     `db:"company_id" json:"company_id"`
	CompanyName     string    `db:"company_name" json:"company_name"`
	CompanyPassword string    `db:"company_password" json:"company_password"`
	CreateTime      time.Time `db:"create_time" json:"create_time"`
	UpdateTime      time.Time `db:"update_time" json:"update_time"`
}
