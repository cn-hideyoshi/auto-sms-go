package mdoel_msg

import (
	"blog.hideyoshi.top/common/pkg/db/model"
	"time"
)

type MsgGroup struct {
	GroupId       int64     `db:"group_id" json:"group_id"`
	GroupName     string    `db:"group_name" json:"group_name"`
	GroupContent  string    `db:"group_content" json:"group_content"`
	GroupType     int32     `db:"group_type" json:"group_type"`
	GroupSendTime time.Time `db:"group_send_time" json:"group_send_time"`
	GroupStatus   int32     `db:"group_status" json:"group_status"`
	CompanyId     int64     `db:"company_id" json:"company_id"`
	TemplateId    int64     `db:"template_id" json:"template_id"`
	model.BaseModel
}

type MsgGroupUser struct {
	GroupId  int64  `db:"group_id" json:"group_id"`
	UserId   int64  `db:"user_id" json:"user_id"`
	UserName string `db:"user_name" json:"user_name"`
	PhoneId  int64  `db:"phone_id" json:"phone_id"`
	PhoneNo  string `db:"phone_no" json:"phone_no"`
}

type MsgTemplate struct {
	TemplateId       int64  `db:"template_id" json:"template_id"`
	TemplateContent  string `db:"template_content" json:"template_content"`
	TemplateCode     string `db:"template_code" json:"template_code"`
	TemplateParamKey string `db:"template_param_key" json:"template_param_key"`
	TemplateSign     string `db:"template_sign" json:"template_sign"`
	model.BaseModel
}
