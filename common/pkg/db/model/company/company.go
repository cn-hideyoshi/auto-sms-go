package model_company

import (
	"blog.hideyoshi.top/common/pkg/db/model"
	"time"
)

type Company struct {
	CompanyId       int64     `db:"company_id" json:"company_id"`
	CompanyName     string    `db:"company_name" json:"company_name"`
	CompanyPassword string    `db:"company_password" json:"company_password"`
	CreateTime      time.Time `db:"create_time" json:"create_time"`
	UpdateTime      time.Time `db:"update_time" json:"update_time"`
}

type Department struct {
	DepartmentId     int64  `db:"department_id" json:"department_id"`
	DepartmentName   string `db:"department_name" json:"department_name"`
	DepartmentParent int64  `db:"department_parent" json:"department_parent"`
	CompanyId        int64  `db:"company_id" json:"company_id"`
	IsRoot           int64  `db:"is_root" json:"is_root"`
	model.BaseModel
}
