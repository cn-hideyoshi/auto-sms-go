package handler

import (
	"blog.hideyoshi.top/common/pkg/db/model"
	"blog.hideyoshi.top/common/pkg/ecode"
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/common/utils"
	"blog.hideyoshi.top/company/internal/cache"
	"blog.hideyoshi.top/company/internal/db/dao"
	"blog.hideyoshi.top/company/pkg/util"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/copier"
	"time"
)

type CompanyLoginHandler struct {
	CompanyHandler
}

func (ch *CompanyLoginHandler) Register(req *companyV1.CompanyRegisterRequest) (*companyV1.CompanyRegisterResponse, error) {
	companyDao := dao.CompanyDao{}

	companyRes := &companyV1.CompanyResponse{
		Code: ecode.SUCCESS,
		Msg:  ecode.GetMsg(ecode.SUCCESS),
	}
	res := &companyV1.CompanyRegisterResponse{
		Response: companyRes,
	}

	info, err := companyDao.GetCompanyByName(req.Username)
	if info != nil && err == nil {
		companyRes.Code = ecode.CompanyExists
		companyRes.Msg = ecode.GetMsg(ecode.CompanyExists)
		return res, nil
	} else if err != nil {
		companyRes.Code = ecode.ERROR
		companyRes.Msg = ecode.GetMsg(ecode.ERROR)
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
		util.SetErrors(companyRes, ecode.ERROR)
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

	claims := jwt.MapClaims{
		"company_name": company.CompanyName,
		"company_id":   company.CompanyId,
		"create_time":  company.CreateTime,
		"update_time":  company.UpdateTime,
	}
	jwtUtils := utils.JWTUtils{
		Claims: claims,
		Method: jwt.SigningMethodES256,
	}
	encode, err := jwtUtils.Encode()
	if err != nil {
		util.SetErrors(companyRes, ecode.ERROR)
		return res, nil
	}

	marshal, err := json.Marshal(company)
	if err != nil {
		util.SetErrors(companyRes, ecode.ERROR)
		return res, nil
	}
	err = cache.Set(fmt.Sprintf("token:%s", encode), string(marshal), 1000)
	if err != nil {
		util.SetErrors(companyRes, ecode.ERROR)
		return res, nil
	}
	res.Token = encode

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
		util.SetErrors(companyRes, ecode.ERROR)
		return &companyV1.CheckCompanyTokenResponse{
			Response: companyRes,
		}
	}

	info := &companyV1.CompanyInfo{}
	copier.Copy(info, decode)

	return &companyV1.CheckCompanyTokenResponse{
		Response:    companyRes,
		CompanyInfo: info,
	}
}
