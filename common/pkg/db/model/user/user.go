package model_user

import "blog.hideyoshi.top/common/pkg/db/model"

type User struct {
	UserId       int64  `db:"user_id" json:"user_id"`
	UserName     string `db:"user_name" json:"user_name"`
	UserPassword string `db:"user_password" json:"user_password"`
	CompanyId    int64  `db:"company_id" json:"company_id"`
	DepartmentId int64  `db:"department_id" json:"department_id"`
	model.BaseModel
}

type UserPhone struct {
	PhoneId int64  `db:"phone_id" json:"phone_id"`
	UserId  int64  `db:"user_id" json:"user_id"`
	PhoneNo string `db:"phone_no" json:"phone_no"`
}
