package crontab

import (
	"blog.hideyoshi.top/msg/internal/amqp"
	"blog.hideyoshi.top/msg/internal/db/dao"
	"encoding/json"
	"log"
	"time"
)

type SMSCrontab struct {
}

func (s SMSCrontab) CheckSmsTime() {
	queue := amqp.SmsQueue{}
	groupDao := dao.MsgGroupDao{}
	groupUserDao := dao.MsgGroupUserDao{}
	templateDao := dao.MsgTemplateDao{}
	group, err := groupDao.SelectTimeToSendGroup(time.Now())
	if err != nil {
		log.Println("SelectTimeToSendGroup err:", err)
	}

	for _, msgGroup := range group {
		users, _ := groupUserDao.SelectMsgGroupByGroupId(msgGroup.GroupId)
		template, _ := templateDao.GetMsgTemplateById(msgGroup.TemplateId)
		body := amqp.SmsBody{
			Info:     msgGroup,
			Users:    users,
			Template: template,
		}
		marshal, _ := json.Marshal(body)

		msgGroup.GroupStatus = 1
		groupDao.UpdateMsgGroup(msgGroup, []string{"group_status"})

		queue.Push(string(marshal))
	}
}
