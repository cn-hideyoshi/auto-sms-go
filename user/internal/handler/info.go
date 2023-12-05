package handler

import (
	model "blog.hideyoshi.top/common/pkg/db/model/user"
	"blog.hideyoshi.top/common/pkg/ecode"
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	userV1 "blog.hideyoshi.top/common/pkg/service/user.v1"
	"blog.hideyoshi.top/user/internal/db/dao"
	"blog.hideyoshi.top/user/pkg/util"
	"blog.hideyoshi.top/user/rpc"
	"database/sql"
	"errors"
	"golang.org/x/net/context"
	"log"
	"time"
)

type InfoHandler struct {
	UserHandler
}

func (h InfoHandler) CreateUserInfo(req *userV1.CreateUserInfoRequest) *userV1.CreateUserInfoResponse {
	res := &userV1.CreateUserInfoResponse{
		Response: &userV1.UserResponse{},
	}

	userDao := dao.UserDao{}
	user := &model.User{
		UserId:       req.UserInfo.UserId,
		UserName:     req.UserInfo.UserName,
		CompanyId:    req.UserInfo.CompanyId,
		DepartmentId: req.UserInfo.DepartmentId,
		UserPassword: req.UserInfo.UserPassword,
	}

	//TODO.. call client use common func
	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	department, err := rpc.Clients.DepartmentInfoClient.GetDepartment(timeout, &companyV1.GetDepartmentRequest{
		DepartmentId: user.DepartmentId,
	})
	if department.Response.Code != 200 {
		util.SetErrors(res.Response, department.Response.Code)
		return res
	}
	if err != nil {
		log.Println("create user info rpc department client err:", err)
		util.SetErrors(res.Response, ecode.ERROR)
		return res
	}

	name, err := userDao.GetUserByName(user.UserName)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("model get user err:", err)
		util.SetErrors(res.Response, ecode.ERROR)
		return res
	}
	if name != nil {
		util.SetErrors(res.Response, ecode.UserExists)
		return res
	}

	err = userDao.CreateUser(user)
	if err != nil {
		log.Println("model create user err:", err)
		util.SetErrors(res.Response, ecode.ERROR)
		return res
	}

	res.UserInfo = &userV1.UserInfo{
		UserId:       user.UserId,
		UserName:     user.UserName,
		UserPassword: user.UserPassword,
		CompanyId:    user.CompanyId,
		DepartmentId: user.DepartmentId,
		CreateTime:   user.CreateTime.Unix(),
		UpdateTime:   user.UpdateTime.Unix(),
	}
	util.SetErrors(res.Response, ecode.SUCCESS)
	return res
}

func (h InfoHandler) UpdateUserInfo(req *userV1.UpdateUserInfoRequest) *userV1.UpdateUserInfoResponse {

	res := &userV1.UpdateUserInfoResponse{
		Response: &userV1.UserResponse{},
	}

	userDao := dao.UserDao{}
	user := &model.User{
		UserId:       req.UserInfo.UserId,
		UserName:     req.UserInfo.UserName,
		UserPassword: req.UserInfo.UserPassword,
		DepartmentId: req.UserInfo.DepartmentId,
		CompanyId:    req.UserInfo.CompanyId,
	}
	updateKey := []string{"user_id", "user_password", "department_id", "company_id"}
	get, err := userDao.GetUserByName(req.UserInfo.UserName)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("update get user name err:", err)
		util.SetErrors(res.Response, ecode.ERROR)
		return res
	} else if get != nil && get.UserId != user.UserId {
		util.SetErrors(res.Response, ecode.UserExists)
		return res
	} else if get == nil {
		updateKey = append(updateKey, "user_name")
	}
	err = userDao.UpdateUser(user, updateKey)
	if err != nil {
		log.Println("update user info err:", err)
		util.SetErrors(res.Response, ecode.ERROR)
		return res
	}
	res.UserInfo = &userV1.UserInfo{
		UserId:       user.UserId,
		UserName:     user.UserName,
		UserPassword: user.UserPassword,
		CompanyId:    user.CompanyId,
		DepartmentId: user.DepartmentId,
		UpdateTime:   user.UpdateTime.Unix(),
	}
	util.SetErrors(res.Response, ecode.SUCCESS)
	return res
}

func (h InfoHandler) GetUserInfo(req *userV1.GetUserInfoRequest) *userV1.GetUserInfoResponse {
	res := &userV1.GetUserInfoResponse{
		Response: &userV1.UserResponse{},
	}
	userDao := dao.UserDao{}
	user, err := userDao.GetUserById(req.UserId)
	if err != nil {
		util.SetErrors(res.Response, ecode.UserNoExists)
		return res
	}

	res.UserInfo = &userV1.UserInfo{
		UserId:     user.UserId,
		UserName:   user.UserName,
		CreateTime: user.CreateTime.Unix(),
		UpdateTime: user.UpdateTime.Unix(),
	}
	util.SetErrors(res.Response, ecode.SUCCESS)
	return res
}
