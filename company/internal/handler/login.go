package handler

import (
	"blog.hideyoshi.top/common/pkg/db/model"
	"blog.hideyoshi.top/common/pkg/ecode"
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/company/internal/db/dao"
	"log"
	"time"
)

type CompanyLoginHandler struct {
	CompanyHandler
}

func (ch *CompanyHandler) Register(req *companyV1.CompanyRegisterRequest) (*companyV1.CompanyRegisterResponse, error) {
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
		companyRes.Code = ecode.COMPANY_EXISTS
		companyRes.Msg = ecode.GetMsg(ecode.COMPANY_EXISTS)
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
		log.Println(err)
		companyRes.Code = ecode.ERROR
		companyRes.Msg = ecode.GetMsg(ecode.ERROR)
	}
	return res, nil
}

func (ch *CompanyHandler) Login(req *companyV1.CompanyLoginRequest) (*companyV1.CompanyLoginResponse, error) {
	return &companyV1.CompanyLoginResponse{
		Response: &companyV1.CompanyResponse{
			Code: ecode.SUCCESS,
			Msg:  ecode.GetMsg(ecode.SUCCESS),
		},
		Token: "test",
	}, nil
}
