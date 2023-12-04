package dao

import (
	"blog.hideyoshi.top/common/pkg/db/model"
	"fmt"
	"strings"
	"time"
)

type UserDao struct {
	field model.User
}

func (c *UserDao) CreateUser(user *model.User) error {
	begin := _db.MustBegin()
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	result, err := begin.NamedExec("insert into as_user (user_id,user_password,user_name,department_id,company_id,create_time,update_time) values (:user_id, :user_password,:user_name,:department_id,:company_id,:create_time,:update_time)", user)
	if err != nil {
		return err
	}

	if err := begin.Commit(); err != nil {
		return err
	}

	if user.UserId, err = result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

func (c *UserDao) UpdateUser(user *model.User, updateKey []string) error {
	begin := _db.MustBegin()

	for i, key := range updateKey {
		updateKey[i] = fmt.Sprintf("%s=:%s", key, key)
	}
	_, err := begin.NamedExec(fmt.Sprintf("update as_user set %s where user_id=:user_id", strings.Join(updateKey, ",")), user)
	if err != nil {
		return err
	}

	if err := begin.Commit(); err != nil {
		return err
	}

	return nil
}

func (c *UserDao) GetUserById(id int64) (*model.User, error) {
	user := &model.User{}
	err := _db.Get(user, "select * from as_user where user_id=? limit 1", id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *UserDao) GetUserByName(name string) (*model.User, error) {
	user := &model.User{}
	err := _db.Get(user, "select * from as_user where user_name=? limit 1", name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *UserDao) SelectUser(company *model.Company) error {
	// TODO: Implement logic for SelectCompany if needed
	return nil
}
