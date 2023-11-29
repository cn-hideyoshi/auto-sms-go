package util

import (
	"blog.hideyoshi.top/common/pkg/ecode"
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/common/utils"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
)

func SetErrors(resp *companyV1.CompanyResponse, code int32) {
	resp.Code = code
	resp.Msg = ecode.GetMsg(code)
}

func GenerateToken(data utils.JwtClaims) (string, error) {
	claims := data
	jwtUtils := utils.JWTUtils{
		Claims: claims,
		Method: jwt.SigningMethodHS256,
	}
	encode, err := jwtUtils.Encode()
	if err != nil {
		return "", err
	}
	return encode, nil
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStr(strLen int) string {
	rand_bytes := make([]rune, strLen)
	for i := range rand_bytes {
		rand_bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(rand_bytes)
}
