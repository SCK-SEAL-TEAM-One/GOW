package repository

import (
	"database/sql"
	"gow/model"
)

type CompanyRepository interface {
	Insert(model.NewCompany) (bool, error)
	GetByTaxID(string) (model.CompanyInfo, error)
}
type CompanyRepositoryMySQL struct {
	DBConnection *sql.DB
}
