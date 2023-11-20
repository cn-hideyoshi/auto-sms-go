package dao

import (
	"blog.hideyoshi.top/common/pkg/db/model"
	"database/sql"
)

type CompanyDao struct {
}

func (c *CompanyDao) CreateCompany(company *model.Company) error {
	begin := _db.MustBegin()
	result, err := begin.NamedExec("insert into wg_company (company_name,company_password,create_time,update_time) values (:company_name, :company_password,:create_time,:update_time)", company)
	if err != nil {
		return err
	}
	if err := begin.Commit(); err != nil {
		return err
	}

	if company.CompanyId, err = result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

func (c *CompanyDao) GetCompanyById(id int64) (*model.Company, error) {
	company := &model.Company{}
	err := _db.Get(company, "select * from wg_company where company_id=? limit 1", id)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (c *CompanyDao) GetCompanyByName(name string) (*model.Company, error) {
	company := &model.Company{}
	err := _db.Get(company, "select * from wg_company where company_name=? limit 1", name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return company, nil
}

func (c *CompanyDao) SelectCompany(company *model.Company) error {
	return nil
}
