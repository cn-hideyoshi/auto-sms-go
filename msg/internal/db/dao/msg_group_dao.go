package dao

import (
	model "blog.hideyoshi.top/common/pkg/db/model/msg"
	"fmt"
	"strings"
	"time"
)

type MsgGroupDao struct {
	field model.MsgGroup
}

func (c *MsgGroupDao) CreateMsgGroup(group *model.MsgGroup) error {
	begin := _db.MustBegin()
	group.CreateTime = time.Now()
	group.UpdateTime = time.Now()
	result, err := begin.NamedExec("insert into as_msg_group (group_id,group_name,group_content,group_type,company_id,template_id,create_time,update_time) values (:group_id,:group_name,:group_content,:group_type,:company_id,:template_id,:create_time,:update_time)", group)
	if err != nil {
		return err
	}

	if err := begin.Commit(); err != nil {
		begin.Rollback()
		return err
	}

	if group.GroupId, err = result.LastInsertId(); err != nil {
		begin.Rollback()
		return err
	}

	return nil
}

func (c *MsgGroupDao) UpdateMsgGroup(user *model.MsgGroup, updateKey []string) error {
	begin := _db.MustBegin()
	updateKey = append(updateKey, "update_time")
	for i, key := range updateKey {
		updateKey[i] = fmt.Sprintf("%s=:%s", key, key)
	}
	user.UpdateTime = time.Now()
	_, err := begin.NamedExec(fmt.Sprintf("update as_msg_group set %s where group_id=:group_id", strings.Join(updateKey, ",")), user)
	if err != nil {
		begin.Rollback()
		return err
	}

	if err := begin.Commit(); err != nil {
		begin.Rollback()
		return err
	}

	return nil
}

func (c *MsgGroupDao) GetMsgGroupById(id int64) (*model.MsgGroup, error) {
	user := &model.MsgGroup{}
	err := _db.Get(user, "select * from as_msg_group where group_id=? limit 1", id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *MsgGroupDao) SelectUser(company *model.MsgGroup) error {
	// TODO: Implement logic for SelectCompany if needed
	return nil
}
