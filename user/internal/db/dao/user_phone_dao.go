package dao

import (
	model "blog.hideyoshi.top/common/pkg/db/model/user"
	"fmt"
	"strings"
)

type UserPhoneDao struct {
	field model.UserPhone
}

func (c *UserPhoneDao) CreateUserPhone(phone *model.UserPhone) error {
	begin := _db.MustBegin()
	result, err := begin.NamedExec("insert into as_user_phone (user_id,phone_no) values (:user_id,:phone_no)", phone)
	if err != nil {
		return err
	}

	if err := begin.Commit(); err != nil {
		return err
	}

	if phone.UserId, err = result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

func (c *UserPhoneDao) UpdateUserPhone(phone *model.UserPhone, updateKey []string) error {
	begin := _db.MustBegin()
	for i, key := range updateKey {
		updateKey[i] = fmt.Sprintf("%s=:%s", key, key)
	}
	_, err := begin.NamedExec(fmt.Sprintf("update as_user_phone set %s where phone_id=:phone_id", strings.Join(updateKey, ",")), phone)
	if err != nil {
		return err
	}

	if err := begin.Commit(); err != nil {
		return err
	}

	return nil
}

func (c *UserPhoneDao) GetUserPhoneById(id int64) (*model.UserPhone, error) {
	phone := &model.UserPhone{}
	err := _db.Get(phone, "select * from as_user_phone where phone_id=? limit 1", id)
	if err != nil {
		return nil, err
	}
	return phone, nil
}

func (c *UserPhoneDao) GetUserPhoneByUserId(id int64) (*model.UserPhone, error) {
	phone := &model.UserPhone{}
	err := _db.Get(phone, "select * from as_user_phone where user_id=? limit 1", id)
	if err != nil {
		return nil, err
	}
	return phone, nil
}

func (c *UserPhoneDao) SelectUserPhone(company *model.UserPhone) error {
	// TODO: Implement logic for SelectCompany if needed
	return nil
}
