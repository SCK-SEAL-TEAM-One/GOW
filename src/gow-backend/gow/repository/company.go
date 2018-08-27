package repository

import (
	"database/sql"
	"gow/model"
	"time"
)

type CompanyRepository interface {
	Insert(model.NewCompany) (bool, error)
	GetByTaxID(string) (model.CompanyInfo, error)
}
type CompanyRepositoryMySQL struct {
	DBConnection *sql.DB
}

func (comrepo CompanyRepositoryMySQL) Insert(newCompany model.NewCompany) (bool, error) {
	statementInsert, err := comrepo.DBConnection.Prepare(`INSERT INTO company 
	(company_name,company_branch,company_address,company_taxid,company_phonenumber,created_time,updated_time)
	VALUES( ?, ?, ?, ?, ?, ?, ? )`)
	if err != nil {
		return false, err
	}
	defer statementInsert.Close()
	_, err = statementInsert.Exec(newCompany.Company, newCompany.Branch, newCompany.Address, newCompany.TaxID, newCompany.PhoneNumber, time.Now(), time.Now())
	if err != nil {
		return false, err
	}
	return true, nil

}
