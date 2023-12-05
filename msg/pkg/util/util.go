package util

import (
	"blog.hideyoshi.top/common/pkg/ecode"
	msgV1 "blog.hideyoshi.top/common/pkg/service/msg.v1"
)

func SetErrors(resp *msgV1.MsgResponse, code int32) {
	resp.Code = code
	resp.Msg = ecode.GetMsg(code)
}
