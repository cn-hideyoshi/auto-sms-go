package model

type User struct {
	UserId       int64  `db:"user_id" json:"user_id"`
	UserName     string `db:"user_name" json:"user_name"`
	UserPassword string `db:"user_password" json:"user_password"`
	CompanyId    int64  `db:"company_id" json:"company_id"`
	DepartmentId int64  `db:"department_id" json:"department_id"`
	BaseModel
}
