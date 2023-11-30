package dao

import (
	"blog.hideyoshi.top/common/pkg/db/model"
	"blog.hideyoshi.top/common/types"
	"fmt"
	"strings"
)

type DepartmentDao struct {
	page *types.Pages
}
type DepartmentDaoOption interface {
	func()
}

func (d *DepartmentDao) CreateDepartment(department *model.Department) error {
	begin := _db.MustBegin()
	result, err := begin.NamedExec("insert into as_department (department_name,department_parent,company_id,create_time,update_time) values (:department_name,:department_parent,:company_id,:create_time,:update_time)", department)
	if err != nil {
		return err
	}

	if err := begin.Commit(); err != nil {
		return err
	}

	if department.DepartmentId, err = result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

func (d *DepartmentDao) UpdateDepartment(department *model.Department, updateKey []string) error {
	begin := _db.MustBegin()

	for i, key := range updateKey {
		updateKey[i] = fmt.Sprintf("%s=:%s", key, key)
	}
	_, err := begin.NamedExec(fmt.Sprintf("update as_department set %s where department_id=:department_id", strings.Join(updateKey, ",")), department)
	if err != nil {
		return err
	}

	if err := begin.Commit(); err != nil {
		return err
	}

	return nil
}

func (d *DepartmentDao) GetDepartmentById(id int64) (*model.Department, error) {
	department := &model.Department{}
	err := _db.Get(department, "select * from as_department where department_id=? limit 1", id)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (d *DepartmentDao) SelectCompany(department *model.Department, opt ...func()) ([]*model.Department, error) {
	for _, f := range opt {
		f()
	}
	departments := make([]*model.Department, 198)
	err := _db.Select(&departments, "select * from as_department where company_id=? limit 999", department.CompanyId)
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (d *DepartmentDao) WithPage(page, pageSize int32) {
	d.page = &types.Pages{
		Page:     page,
		PageSize: pageSize,
	}
}
