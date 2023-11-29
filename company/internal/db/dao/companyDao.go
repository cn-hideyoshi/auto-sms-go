package dao

import (
	"blog.hideyoshi.top/common/pkg/db/model"
	"fmt"
	"strings"
)

// CompanyDao represents the data access object for the Company model.
type CompanyDao struct {
	field model.Company
}

// CreateCompany inserts a new company record into the database.
func (c *CompanyDao) CreateCompany(company *model.Company) error {
	// Start a transaction
	begin := _db.MustBegin()

	// Execute the insert query using named parameters
	result, err := begin.NamedExec("insert into as_company (company_name,company_password,create_time,update_time) values (:company_name, :company_password,:create_time,:update_time)", company)
	if err != nil {
		return err
	}

	// Commit the transaction
	if err := begin.Commit(); err != nil {
		return err
	}

	// Retrieve the last inserted ID and update the company object
	if company.CompanyId, err = result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

// UpdateCompany updates a company record in the database based on the given update keys.
func (c *CompanyDao) UpdateCompany(company *model.Company, updateKey []string) error {
	// Start a transaction
	begin := _db.MustBegin()

	// Build the update query with named parameters
	for i, key := range updateKey {
		updateKey[i] = fmt.Sprintf("%s=:%s", key, key)
	}
	_, err := begin.NamedExec(fmt.Sprintf("update as_company set %s where company_id=:company_id", strings.Join(updateKey, ",")), company)
	if err != nil {
		return err
	}

	// Commit the transaction
	if err := begin.Commit(); err != nil {
		return err
	}

	return nil
}

// GetCompanyById retrieves a company record from the database based on the company ID.
func (c *CompanyDao) GetCompanyById(id int64) (*model.Company, error) {
	company := &model.Company{}
	err := _db.Get(company, "select * from as_company where company_id=? limit 1", id)
	if err != nil {
		return nil, err
	}
	return company, nil
}

// GetCompanyByName retrieves a company record from the database based on the company name.
func (c *CompanyDao) GetCompanyByName(name string) (*model.Company, error) {
	company := &model.Company{}
	err := _db.Get(company, "select * from as_company where company_name=? limit 1", name)
	if err != nil {
		return nil, err
	}
	return company, nil
}

// SelectCompany is a placeholder function without implementation. Add logic as needed.
func (c *CompanyDao) SelectCompany(company *model.Company) error {
	// TODO: Implement logic for SelectCompany if needed
	return nil
}
