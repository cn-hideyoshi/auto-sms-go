package handler

import (
	model "blog.hideyoshi.top/common/pkg/db/model/user"
	"blog.hideyoshi.top/common/pkg/ecode"
	userV1 "blog.hideyoshi.top/common/pkg/service/user.v1"
	"blog.hideyoshi.top/user/internal/db/dao"
	"blog.hideyoshi.top/user/pkg/util"
	"database/sql"
	"errors"
	"log"
)

type PhoneHandler struct {
	UserHandler
}

func (h PhoneHandler) CreateUserPhone(req *userV1.CreateUserPhoneRequest) *userV1.CreateUserPhoneResponse {
	var err error
	d := dao.UserPhoneDao{}
	res := &userV1.CreateUserPhoneResponse{
		Response: &userV1.UserResponse{},
	}
	phone := &model.UserPhone{
		UserId:  req.UserId,
		PhoneNo: req.PhoneNo,
	}

	get, err := d.GetUserPhoneByUserId(req.UserId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("create phone-> get user phone by id err:", err)
		util.SetErrors(res.Response, ecode.ERROR)
		return res
	}
	if get != nil {
		phone.PhoneId = get.PhoneId
		err = d.UpdateUserPhone(phone, []string{"phone_no"})
	} else {
		err = d.CreateUserPhone(phone)
	}
	if err != nil {
		log.Println("create user phone err:", err)
		util.SetErrors(res.Response, ecode.ERROR)
	}
	res.UserPhone = &userV1.UserPhone{
		PhoneId: phone.PhoneId,
		UserId:  phone.UserId,
		PhoneNo: phone.PhoneNo,
	}
	util.SetErrors(res.Response, ecode.SUCCESS)
	return res
}

func (h PhoneHandler) GetUserPhone(req *userV1.GetUserPhoneRequest) *userV1.GetUserPhoneResponse {
	d := dao.UserPhoneDao{}
	res := &userV1.GetUserPhoneResponse{
		Response: &userV1.UserResponse{},
	}

	get, err := d.GetUserPhoneByUserId(req.UserId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("create phone-> get user phone by id err:", err)
		util.SetErrors(res.Response, ecode.ERROR)
		return res
	} else if get == nil {
		util.SetErrors(res.Response, ecode.UserNoExists)
		return res
	}
	res.UserPhone = &userV1.UserPhone{
		UserId:  get.UserId,
		PhoneNo: get.PhoneNo,
		PhoneId: get.PhoneId,
	}
	util.SetErrors(res.Response, ecode.SUCCESS)
	return res
}
