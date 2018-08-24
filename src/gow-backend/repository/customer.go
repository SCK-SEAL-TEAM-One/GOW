package repository

import (
	"database/sql"
	"gow-backend/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepository interface {
	Insert(model.NewCustomer) (bool, error)
	GetByTaxID(string) (model.CustomerInfo, error)
}
type CustomerRepositoryMySQL struct {
	ConnetionDB *sql.DB
}

func (repository CustomerRepositoryMySQL) Insert(newCustomer model.NewCustomer) (bool, error) {
	statementInsert, err := repository.ConnetionDB.Prepare(`INSERT INTO customer 
	(customer_name,customer_branch,customer_address,customer_taxid,created_time,updated_time)
	VALUES( ?, ?, ?, ?, ?, ? )`)
	if err != nil {
		return false, err
	}
	defer statementInsert.Close()
	_, err = statementInsert.Exec(newCustomer.Company, newCustomer.Branch, newCustomer.Address, newCustomer.TaxID, time.Now(), time.Now())
	if err != nil {
		return false, err
	}
	return true, nil
}
func (repository CustomerRepositoryMySQL) GetByTaxID(string) (model.CustomerInfo, error) {
	return model.CustomerInfo{
		ID:      1,
		Company: "บริษัท ที.เอ็น. อินคอร์ปอเรชั่นจำกัด",
		Branch:  "สำนักงานใหญ่",
		Address: "3 อาคารรัจนาการ ถนนสาทรใต้ แขวงยานนาวา เขตสาทร กรุงเทพมหานคร 10120",
		TaxID:   "0105553108372",
	}, nil
}
