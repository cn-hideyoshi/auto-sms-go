package dao

import (
	model "blog.hideyoshi.top/common/pkg/db/model/msg"
	"fmt"
	"strings"
)

type MsgGroupUserDao struct {
	field model.MsgGroup
}

func (c *MsgGroupUserDao) CreateMsgGroupUser(user *model.MsgGroupUser) error {
	begin := _db.MustBegin()
	result, err := begin.NamedExec("insert into as_msg_group_user (group_id,user_id,phone_id,phone_no) values (:group_id,:user_id,:phone_id,:phone_no)", user)
	if err != nil {
		return err
	}

	if err := begin.Commit(); err != nil {
		return err
	}

	if user.GroupId, err = result.LastInsertId(); err != nil {
		return err
	}

	return nil
}
func (c *MsgGroupUserDao) BatchCreateMsgGroupUser(users []*model.MsgGroupUser) error {
	begin := _db.MustBegin()
	for _, user := range users {
		_, err := begin.NamedExec("insert into as_msg_group_user (group_id,user_id,phone_id,phone_no) values (:group_id,:user_id,:phone_id,:phone_no)", user)
		if err != nil {
			begin.Rollback()
			return err
		}
		//if user.GroupId, err = result.LastInsertId(); err != nil {
		//	begin.Rollback()
		//	return err
		//}
	}
	if err := begin.Commit(); err != nil {
		return err
	}
	return nil
}
func (c *MsgGroupUserDao) UpdateMsgGroupUser(user *model.MsgGroupUser, updateKey []string) error {
	begin := _db.MustBegin()
	updateKey = append(updateKey, "update_time")
	for i, key := range updateKey {
		updateKey[i] = fmt.Sprintf("%s=:%s", key, key)
	}
	_, err := begin.NamedExec(fmt.Sprintf("update as_msg_group_user set %s where group_id=:group_id", strings.Join(updateKey, ",")), user)
	if err != nil {
		return err
	}

	if err := begin.Commit(); err != nil {
		return err
	}

	return nil
}

func (c *MsgGroupUserDao) GetMsgGroupByGroupId(id int64) (*model.MsgGroupUser, error) {
	user := &model.MsgGroupUser{}
	err := _db.Get(user, "select * from as_msg_group_user where group_id=? limit 1", id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *MsgGroupUserDao) SelectMsgGroupByGroupId(id int64) ([]*model.MsgGroupUser, error) {
	departments := make([]*model.MsgGroupUser, 10)
	err := _db.Select(&departments, "select * from as_msg_group_user where group_id=? limit 999", id)
	if err != nil {
		return nil, err
	}
	return departments, nil
}
