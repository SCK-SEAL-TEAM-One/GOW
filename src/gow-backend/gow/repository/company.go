package repository

import (
	"database/sql"
	"gow/model"
	"time"
)

type CompanyRepository interface {
	Insert(model.NewCompany) (bool, error)
	GetByTaxID(string) (model.CompanyInfo, error)
	GetAll() ([]model.CompanyInfo, error)
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

func (comrepo CompanyRepositoryMySQL) GetByTaxID(companyTaxID string) (model.CompanyInfo, error) {
	var companyInfo model.CompanyInfo
	statementQuery := `SELECT * FROM company WHERE company_taxid=?`
	row := comrepo.DBConnection.QueryRow(statementQuery, companyTaxID)
	err := row.Scan(
		&companyInfo.ID,
		&companyInfo.Company,
		&companyInfo.Branch,
		&companyInfo.Address,
		&companyInfo.TaxID,
		&companyInfo.PhoneNumber,
		&companyInfo.CreatedTime,
		&companyInfo.UpdatedTime,
	)
	if err != nil {
		return model.CompanyInfo{}, err
	}
	return companyInfo, nil
}

func (comrepo CompanyRepositoryMySQL) GetAll() ([]model.CompanyInfo, error) {
	return []model.CompanyInfo{}, nil
}
