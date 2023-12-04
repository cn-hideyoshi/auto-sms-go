package util

import (
	"blog.hideyoshi.top/common/pkg/ecode"
	userV1 "blog.hideyoshi.top/common/pkg/service/user.v1"
)

func SetErrors(resp *userV1.UserResponse, code int32) {
	resp.Code = code
	resp.Msg = ecode.GetMsg(code)
}
