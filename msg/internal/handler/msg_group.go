package handler

import (
	model "blog.hideyoshi.top/common/pkg/db/model/msg"
	"blog.hideyoshi.top/common/pkg/ecode"
	msgV1 "blog.hideyoshi.top/common/pkg/service/msg.v1"
	"blog.hideyoshi.top/msg/internal/db/dao"
	"blog.hideyoshi.top/msg/pkg/util"
	"database/sql"
	"errors"
	"log"
	"time"
)

type MsgGroupHandler struct {
	MsgHandler
}

func (h MsgGroupHandler) CreateMsgGroup(req *msgV1.CreateMessageGroupRequest) *msgV1.CreateMessageGroupResponse {
	mDao := dao.MsgGroupDao{}
	muDao := dao.MsgGroupUserDao{}
	var err error
	res := &msgV1.CreateMessageGroupResponse{
		Response: &msgV1.MsgResponse{},
	}

	reqMessageGroup := req.MessageGroup
	if err != nil {
		log.Println("parse time err:", err)
		return nil
	}
	group := &model.MsgGroup{
		CompanyId:     reqMessageGroup.CompanyId,
		GroupName:     reqMessageGroup.GroupName,
		GroupContent:  reqMessageGroup.GroupContent,
		GroupType:     reqMessageGroup.GroupType,
		TemplateId:    reqMessageGroup.TemplateId,
		GroupSendTime: time.Unix(reqMessageGroup.GroupSendTime, 0),
	}
	err = mDao.CreateMsgGroup(group)
	if err != nil {
		log.Println("batch create msg user err:", err)
		util.SetErrors(res.Response, ecode.ERROR)
		return res
	}

	var mus = make([]*model.MsgGroupUser, len(reqMessageGroup.MessageGroupUser))
	for i, user := range reqMessageGroup.MessageGroupUser {
		user.GroupId = group.GroupId
		mus[i] = &model.MsgGroupUser{
			GroupId: group.GroupId,
			UserId:  user.UserId,
			PhoneId: user.PhoneId,
			PhoneNo: user.PhoneNo,
		}
	}

	err = muDao.BatchCreateMsgGroupUser(mus)
	if err != nil {
		log.Println("batch create msg user err:", err)
		util.SetErrors(res.Response, ecode.ERROR)
		return res
	}
	util.SetErrors(res.Response, ecode.SUCCESS)
	res.MessageGroup = &msgV1.MessageGroup{
		GroupId:          group.GroupId,
		CompanyId:        group.CompanyId,
		TemplateId:       group.TemplateId,
		GroupName:        group.GroupName,
		GroupType:        group.GroupType,
		GroupContent:     group.GroupContent,
		CreateTime:       group.CreateTime.Unix(),
		UpdateTime:       group.UpdateTime.Unix(),
		MessageGroupUser: reqMessageGroup.MessageGroupUser,
	}
	return res
}

func (h MsgGroupHandler) UpdateMsgGroup(req *msgV1.UpdateMessageGroupRequest) *msgV1.UpdateMessageGroupResponse {
	res := &msgV1.UpdateMessageGroupResponse{
		Response: &msgV1.MsgResponse{},
	}
	return res
}

func (h MsgGroupHandler) GetMsgGroup(req *msgV1.GetMessageGroupRequest) *msgV1.GetMessageGroupResponse {
	res := &msgV1.GetMessageGroupResponse{
		Response: &msgV1.MsgResponse{},
	}
	mDao := dao.MsgGroupDao{}
	muDao := dao.MsgGroupUserDao{}
	var err error
	get, err := mDao.GetMsgGroupById(req.GroupId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		util.SetErrors(res.Response, ecode.MsgGroupNoExists)
		return res
	} else if err != nil {
		log.Println("get msg err:", err)
		util.SetErrors(res.Response, ecode.MsgGroupNoExists)
		return res
	}

	getUsers, err := muDao.SelectMsgGroupByGroupId(get.GroupId)
	users := make([]*msgV1.MessageGroupUser, len(getUsers))
	for i, user := range getUsers {
		users[i] = &msgV1.MessageGroupUser{
			UserId:  user.UserId,
			GroupId: user.GroupId,
			PhoneNo: user.PhoneNo,
			PhoneId: user.PhoneId,
		}
	}
	res.MessageGroup = &msgV1.MessageGroup{
		GroupId:          get.GroupId,
		GroupName:        get.GroupName,
		GroupContent:     get.GroupContent,
		GroupType:        get.GroupType,
		CompanyId:        get.CompanyId,
		TemplateId:       get.TemplateId,
		CreateTime:       get.CreateTime.Unix(),
		UpdateTime:       get.UpdateTime.Unix(),
		MessageGroupUser: users,
	}

	return res
}
