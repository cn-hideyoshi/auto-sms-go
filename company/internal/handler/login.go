package handler

import (
	model "blog.hideyoshi.top/common/pkg/db/model/company"
	"blog.hideyoshi.top/common/pkg/ecode"
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/common/utils"
	"blog.hideyoshi.top/company/config"
	"blog.hideyoshi.top/company/internal/cache"
	"blog.hideyoshi.top/company/internal/db/dao"
	"blog.hideyoshi.top/company/pkg/util"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

type CompanyLoginHandler struct {
	CompanyHandler
}

func (ch *CompanyLoginHandler) Register(req *companyV1.CompanyRegisterRequest) (*companyV1.CompanyRegisterResponse, error) {
	companyDao := dao.CompanyDao{}
	res := &companyV1.CompanyRegisterResponse{
		Response: &companyV1.CompanyResponse{
			Code: ecode.SUCCESS,
			Msg:  ecode.GetMsg(ecode.SUCCESS),
		},
	}

	info, err := companyDao.GetCompanyByName(req.Username)
	if !errors.Is(err, sql.ErrNoRows) {
		util.SetErrors(res.Response, ecode.ERROR)
		return res, nil
	} else if info != nil {
		util.SetErrors(res.Response, ecode.CompanyExists)
		return res, nil
	}
	company := model.Company{
		CompanyName:     req.Username,
		CompanyPassword: req.Password,
		CreateTime:      time.Now(),
		UpdateTime:      time.Now(),
	}

	err = companyDao.CreateCompany(&company)
	if err != nil {
		util.SetErrors(res.Response, ecode.ERROR)
	}
	return res, nil
}
func (ch *CompanyLoginHandler) Login(req *companyV1.CompanyLoginRequest) (*companyV1.CompanyLoginResponse, error) {
	companyDao := dao.CompanyDao{}
	companyRes := &companyV1.CompanyResponse{
		Code: ecode.SUCCESS,
		Msg:  ecode.GetMsg(ecode.SUCCESS),
	}
	res := &companyV1.CompanyLoginResponse{
		Response: companyRes,
	}

	company, err := companyDao.GetCompanyByName(req.Username)
	if err != nil {
		util.SetErrors(companyRes, ecode.ERROR)
		return res, nil
	}
	jwtClaims := utils.JwtClaims{
		Data: map[string]interface{}{
			"company_id":   company.CompanyId,
			"company_name": company.CompanyName,
			"create_time":  company.CreateTime.Unix(),
			"update_time":  company.UpdateTime.Unix(),
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.C.Grpc.Name + "_server",                // 签发者
			Subject:   company.CompanyName,                           // 签发对象
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                //签发时间
			ID:        util.RandStr(10),                              // wt ID, 类似于盐值
		},
	}
	token, err := util.GenerateToken(jwtClaims)

	if err != nil {
		util.SetErrors(companyRes, ecode.JwtErr)
		return res, nil
	}
	marshal, err := json.Marshal(company)
	if err != nil {
		util.SetErrors(companyRes, ecode.ERROR)
		return res, nil
	}
	err = cache.Cache.Set(fmt.Sprintf("token:%s", token), string(marshal), 1000)
	if err != nil {
		util.SetErrors(companyRes, ecode.RedisErr)
		return res, nil
	}
	res.Token = token

	return res, nil
}

func (ch *CompanyLoginHandler) CheckCompanyToken(req *companyV1.CheckCompanyTokenRequest) *companyV1.CheckCompanyTokenResponse {
	companyRes := &companyV1.CompanyResponse{
		Code: ecode.SUCCESS,
		Msg:  ecode.GetMsg(ecode.SUCCESS),
	}
	jwtUtils := utils.JWTUtils{}
	decode, err := jwtUtils.Decode(req.Token)
	if err != nil {
		log.Println("CheckCompanyToken:line 118,", err)
		util.SetErrors(companyRes, ecode.ERROR)
		return &companyV1.CheckCompanyTokenResponse{
			Response: companyRes,
		}
	}
	data := decode.Data.(map[string]interface{})
	info := &companyV1.CompanyInfo{
		CompanyId:   int64(data["company_id"].(float64)),
		CompanyName: data["company_name"].(string),
		CreateTime:  int64(data["create_time"].(float64)),
		UpdateTime:  int64(data["update_time"].(float64)),
	}
	return &companyV1.CheckCompanyTokenResponse{
		Response:    companyRes,
		CompanyInfo: info,
	}
}
