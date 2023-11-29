package handler

import (
	"blog.hideyoshi.top/common/pkg/db/model"
	"blog.hideyoshi.top/common/pkg/ecode"
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/company/internal/db/dao"
	"blog.hideyoshi.top/company/pkg/util"
	"database/sql"
	"errors"
	"github.com/jinzhu/copier"
	"log"
	"time"
)

// CompanyInfoHandler handles company information related requests.
type CompanyInfoHandler struct {
	CompanyHandler
}

// GetCompanyInfo retrieves information for a specific company based on the provided request.
func (ch *CompanyInfoHandler) GetCompanyInfo(req *companyV1.GetCompanyInfoRequest) (*companyV1.CompanyInfoResponse, error) {
	// Create a CompanyDao instance
	companyDao := dao.CompanyDao{}
	info := &companyV1.CompanyInfo{}
	res := &companyV1.CompanyInfoResponse{
		Response: &companyV1.CompanyResponse{
			Code: ecode.SUCCESS,
			Msg:  ecode.GetMsg(ecode.SUCCESS),
		},
		CompanyInfo: info,
	}

	// Retrieve company information by ID
	company, err := companyDao.GetCompanyById(int64(req.CompanyId))
	if err != nil {
		// Set error response if there is an error
		util.SetErrors(res.Response, ecode.ERROR)
		return res, nil
	}

	// Populate response with company information
	info.CompanyId = company.CompanyId
	info.CompanyName = company.CompanyName
	info.CreateTime = company.CreateTime.Unix()
	info.UpdateTime = company.UpdateTime.Unix()
	return res, nil
}

// UpdateCompanyInfo updates the information for a specific company based on the provided request.
func (ch *CompanyInfoHandler) UpdateCompanyInfo(req *companyV1.UpdateCompanyInfoRequest) (*companyV1.CompanyInfoResponse, error) {
	// Create a CompanyDao instance
	companyDao := dao.CompanyDao{}
	getCompanyInfoRequest := &companyV1.GetCompanyInfoRequest{
		CompanyId: req.CompanyId,
	}
	res, _ := ch.GetCompanyInfo(getCompanyInfoRequest)
	if res.Response.Code != 200 {
		return res, nil
	}

	// Create a response structure
	res = &companyV1.CompanyInfoResponse{
		Response: &companyV1.CompanyResponse{
			Code: ecode.SUCCESS,
			Msg:  ecode.GetMsg(ecode.SUCCESS),
		},
		CompanyInfo: &companyV1.CompanyInfo{},
	}

	// Check if a company with the new name already exists
	getCompany, err := companyDao.GetCompanyByName(req.CompanyInfo.CompanyName)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("update company info err:", err)
		util.SetErrors(res.Response, ecode.ERROR)
		return res, nil
	} else if getCompany != nil && getCompany.CompanyId != req.CompanyId {
		util.SetErrors(res.Response, ecode.CompanyExists)
		return res, nil
	}

	// Copy information from request to the company model
	req.CompanyInfo.CompanyId = req.CompanyId
	company := &model.Company{}
	copier.Copy(company, req.CompanyInfo)
	company.UpdateTime = time.Now()
	updateKey := []string{"company_password", "update_time"}
	if getCompany == nil {
		updateKey = append(updateKey, "company_name")
	}

	// Update the company information in the database
	err = companyDao.UpdateCompany(company, updateKey)
	if err != nil {
		log.Println("update company info err:", err)
		util.SetErrors(res.Response, ecode.ERROR)
		return res, nil
	}

	// Update the response with the updated company information
	res.CompanyInfo.CompanyId = company.CompanyId
	res.CompanyInfo.CompanyName = company.CompanyName
	res.CompanyInfo.UpdateTime = company.UpdateTime.Unix()
	res.CompanyInfo.CreateTime = getCompany.CreateTime.Unix()
	return res, nil
}
