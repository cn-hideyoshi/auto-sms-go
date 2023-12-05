package handler

import (
	model "blog.hideyoshi.top/common/pkg/db/model/company"
	"blog.hideyoshi.top/common/pkg/ecode"
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/company/internal/cache"
	"blog.hideyoshi.top/company/internal/db/dao"
	"blog.hideyoshi.top/company/pkg/util"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
)

type DepartmentInfoHandler struct {
	CompanyHandler
}

func (h DepartmentInfoHandler) CreateDepartment(req *companyV1.CreateDepartmentRequest) *companyV1.CreateDepartmentResponse {
	departmentDao := dao.DepartmentDao{}

	res := &companyV1.CreateDepartmentResponse{
		Response: &companyV1.CompanyResponse{
			Code: ecode.SUCCESS,
			Msg:  ecode.GetMsg(ecode.SUCCESS),
		},
	}
	department := &model.Department{
		DepartmentName:   req.DepartmentInfo.DepartmentName,
		DepartmentParent: req.DepartmentInfo.DepartmentParent,
		CompanyId:        req.DepartmentInfo.CompanyId,
	}
	department.CreateTime = time.Now()
	department.UpdateTime = time.Now()
	err := departmentDao.CreateDepartment(department)
	if err != nil {
		util.SetErrors(res.Response, ecode.ERROR)
		return res
	}
	return res
}

func (h DepartmentInfoHandler) GetCompanyInfo(req *companyV1.GetDepartmentRequest) *companyV1.GetDepartmentResponse {
	defer func() {
		r := recover()
		if r != nil {
			log.Println("get company info panic", r)
		}
	}()
	res := &companyV1.GetDepartmentResponse{
		Response: &companyV1.CompanyResponse{
			Code: ecode.SUCCESS,
			Msg:  ecode.GetMsg(ecode.SUCCESS),
		},
	}

	departmentDao := dao.DepartmentDao{}

	department, err := departmentDao.GetDepartmentById(req.DepartmentId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("get department by id err:", err)
		util.SetErrors(res.Response, ecode.ERROR)
		return res
	}

	res.DepartmentInfo = &companyV1.DepartmentInfo{
		DepartmentId:     department.CompanyId,
		DepartmentName:   department.DepartmentName,
		DepartmentParent: department.DepartmentParent,
		CompanyId:        department.CompanyId,
		CreateTime:       department.CreateTime.Unix(),
		UpdateTime:       department.UpdateTime.Unix(),
	}
	return res
}

func (h DepartmentInfoHandler) UpdateDepartment(req *companyV1.UpdateDepartmentRequest) *companyV1.UpdateDepartmentResponse {
	defer func() {
		r := recover()
		if r != nil {
			log.Println("update company info panic", r)
		}
	}()

	res := &companyV1.UpdateDepartmentResponse{
		Response: &companyV1.CompanyResponse{
			Code: ecode.SUCCESS,
			Msg:  ecode.GetMsg(ecode.SUCCESS),
		},
	}

	department := &model.Department{
		DepartmentId:     req.DepartmentId,
		CompanyId:        req.DepartmentInfo.CompanyId,
		DepartmentName:   req.DepartmentInfo.DepartmentName,
		DepartmentParent: req.DepartmentInfo.DepartmentParent,
	}
	department.UpdateTime = time.Now()
	departmentDao := dao.DepartmentDao{}
	err := departmentDao.UpdateDepartment(department, []string{"company_id", "department_name", "department_parent", "update_time"})
	if err != nil {
		log.Println("update department info err:", err)
		util.SetErrors(res.Response, ecode.ERROR)
		return res
	}
	return res
}

func (h DepartmentInfoHandler) GetDepartmentTree(req *companyV1.GetDepartmentTreeRequest) *companyV1.GetDepartmentTreeResponse {
	res := &companyV1.GetDepartmentTreeResponse{
		Response: &companyV1.CompanyResponse{
			Code: ecode.SUCCESS,
			Msg:  ecode.GetMsg(ecode.SUCCESS),
		},
	}
	redisKey := fmt.Sprintf("department:parent%d:", req.DepartmentParent)
	treeStr, err := cache.Cache.Get(redisKey)
	if err == nil {
		var tree []*companyV1.DepartmentTree
		err := json.Unmarshal([]byte(treeStr), &tree)
		if err == nil {
			res.DepartmentTree = tree
			return res
		}
		log.Println("get department json unmarshal err:", err, "use uncached data")
	}

	departmentDao := dao.DepartmentDao{}
	department := &model.Department{
		CompanyId: req.CompanyId,
	}
	departments, err := departmentDao.SelectCompany(department)
	if err != nil {
		log.Println("get department tree err:", err)
		util.SetErrors(res.Response, ecode.ERROR)
		return res
	}
	tree := h.buildTree(departments, req.DepartmentParent)
	marshal, err := json.Marshal(tree)
	err = cache.Cache.Set(redisKey, string(marshal), 10000)
	res.DepartmentTree = tree
	return res
}

func (h DepartmentInfoHandler) buildTree(departments []*model.Department, parentId int64) []*companyV1.DepartmentTree {
	var tree []*companyV1.DepartmentTree
	for _, department := range departments {
		if department.DepartmentParent == parentId {
			newChild := &companyV1.DepartmentTree{
				DepartmentInfo: &companyV1.DepartmentInfo{
					DepartmentId:     department.DepartmentId,
					DepartmentName:   department.DepartmentName,
					DepartmentParent: department.DepartmentParent,
					CompanyId:        department.CompanyId,
					CreateTime:       department.CreateTime.Unix(),
					UpdateTime:       department.UpdateTime.Unix(),
				},
				Node: h.buildTree(departments, department.DepartmentId),
			}
			tree = append(tree, newChild)
		}
	}
	return tree
}
