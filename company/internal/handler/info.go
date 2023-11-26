package handler

import (
	"blog.hideyoshi.top/common/pkg/ecode"
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/company/internal/db/dao"
	"blog.hideyoshi.top/company/pkg/util"
	"github.com/jinzhu/copier"
)

type CompanyInfoHandler struct {
	CompanyHandler
}

func (ch *CompanyInfoHandler) GetCompanyInfo(req *companyV1.CompanyInfo) (*companyV1.CompanyInfoResponse, error) {
	companyDao := dao.CompanyDao{}

	companyRes := &companyV1.CompanyResponse{
		Code: ecode.SUCCESS,
		Msg:  ecode.GetMsg(ecode.SUCCESS),
	}

	info := &companyV1.CompanyInfo{}
	res := &companyV1.CompanyInfoResponse{
		Response:    companyRes,
		CompanyInfo: info,
	}
	company, err := companyDao.GetCompanyById(int64(req.CompanyId))
	if err != nil {
		util.SetErrors(companyRes, ecode.ERROR)
		return res, nil
	}
	err = copier.Copy(&info, &company)
	if err != nil {
		util.SetErrors(companyRes, ecode.ERROR)
		return res, nil
	}
	return res, nil
}

func (ch *CompanyInfoHandler) UpdateCompanyInfo(req *companyV1.CompanyInfo) (*companyV1.CompanyInfoResponse, error) {
	companyDao := dao.CompanyDao{}

	companyRes := &companyV1.CompanyResponse{
		Code: ecode.SUCCESS,
		Msg:  ecode.GetMsg(ecode.SUCCESS),
	}
	info := &companyV1.CompanyInfo{}
	res := &companyV1.CompanyInfoResponse{
		Response: companyRes,
	}

	company, err := companyDao.GetCompanyById(int64(req.CompanyId))
	if err != nil {
		util.SetErrors(companyRes, ecode.ERROR)
		return res, nil
	}
	err = copier.Copy(&info, &company)
	if err != nil {
		util.SetErrors(companyRes, ecode.ERROR)
		return res, nil
	}
	return res, nil
}
