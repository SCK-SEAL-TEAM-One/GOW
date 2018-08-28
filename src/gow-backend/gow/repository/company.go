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
	company := newCompany.ToCompanyModel()
	statement := `INSERT INTO company 
	(company_name_th,company_name_en,company_branch_th,company_branch_en,company_address_th,company_address_en,
	company_taxid,company_phonenumber,created_time,updated_time)
	VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`
	statementInsert, err := comrepo.DBConnection.Prepare(statement)
	if err != nil {
		return false, err
	}
	defer statementInsert.Close()
	_, err = statementInsert.Exec(company.NameTH, company.NameEN, company.BranchTH, company.BranchEN, company.AddressTH, company.AddressEN, company.TaxID, company.PhoneNumber, time.Now(), time.Now())
	if err != nil {
		return false, err
	}
	return true, nil

}

func (comrepo CompanyRepositoryMySQL) GetByTaxID(companyTaxID string) (model.CompanyInfo, error) {
	var company model.Company
	statementQuery := `SELECT * FROM company WHERE company_taxid=?`
	row := comrepo.DBConnection.QueryRow(statementQuery, companyTaxID)
	err := row.Scan(
		&company.ID,
		&company.NameTH,
		&company.NameEN,
		&company.BranchTH,
		&company.BranchEN,
		&company.AddressTH,
		&company.AddressEN,
		&company.TaxID,
		&company.PhoneNumber,
		&company.CreatedTime,
		&company.UpdatedTime,
	)
	if err != nil {
		return model.CompanyInfo{}, err
	}
	return company.ToCompanyInfo(), nil
}

func (comrepo CompanyRepositoryMySQL) GetAll() ([]model.CompanyInfo, error) {
	var companyInfo []model.CompanyInfo
	statementQuery := `SELECT * FROM company`
	rows, err := comrepo.DBConnection.Query(statementQuery)
	if err != nil {
		return companyInfo, err
	}
	for rows.Next() {
		var company model.Company
		rows.Scan(
			&company.ID,
			&company.NameTH,
			&company.NameEN,
			&company.BranchTH,
			&company.BranchEN,
			&company.AddressTH,
			&company.AddressEN,
			&company.TaxID,
			&company.PhoneNumber,
			&company.CreatedTime,
			&company.UpdatedTime,
		)
		companyInfo = append(companyInfo, company.ToCompanyInfo())
	}
	return companyInfo, nil
}
