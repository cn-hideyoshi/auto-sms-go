package dao

import (
	model "blog.hideyoshi.top/common/pkg/db/model/msg"
)

type MsgTemplateDao struct {
	field model.MsgTemplate
}

func (c *MsgTemplateDao) GetMsgTemplateById(id int64) (*model.MsgTemplate, error) {
	template := &model.MsgTemplate{}
	err := _db.Get(template, "select * from as_msg_template where template_id=? limit 1", id)
	if err != nil {
		return nil, err
	}
	return template, nil
}
