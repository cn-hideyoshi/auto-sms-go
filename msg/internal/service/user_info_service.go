package service

import (
	msgV1 "blog.hideyoshi.top/common/pkg/service/msg.v1"
	"blog.hideyoshi.top/msg/internal/handler"
	"context"
)

type MsgGroupService struct {
	msgV1.UnimplementedMsgGroupServiceServer
}

func NewMsgGroupService() *MsgGroupService {
	return &MsgGroupService{}
}

func (*MsgGroupService) CreateMsgGroup(ctx context.Context, req *msgV1.CreateMessageGroupRequest) (*msgV1.CreateMessageGroupResponse, error) {
	h := handler.MsgGroupHandler{}
	return h.CreateMsgGroup(req), nil
}

func (*MsgGroupService) UpdateMsgGroup(ctx context.Context, req *msgV1.UpdateMessageGroupRequest) (*msgV1.UpdateMessageGroupResponse, error) {
	h := handler.MsgGroupHandler{}
	return h.UpdateMsgGroup(req), nil
}

func (*MsgGroupService) GetMsgGroup(ctx context.Context, req *msgV1.GetMessageGroupRequest) (*msgV1.GetMessageGroupResponse, error) {
	h := handler.MsgGroupHandler{}
	return h.GetMsgGroup(req), nil
}
