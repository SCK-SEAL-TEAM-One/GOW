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
func (repository CustomerRepositoryMySQL) GetByTaxID(customerTaxID string) (model.CustomerInfo, error) {
	var customerInfo model.CustomerInfo
	statementQuery := `SELECT * FROM customer WHERE customer_taxid=?`
	row := repository.ConnetionDB.QueryRow(statementQuery, customerTaxID)
	err := row.Scan(
		&customerInfo.ID,
		&customerInfo.Company,
		&customerInfo.Branch,
		&customerInfo.Address,
		&customerInfo.TaxID,
		&customerInfo.CreatedTime,
		&customerInfo.UpdatedTime,
	)
	if err != nil {
		return model.CustomerInfo{}, err
	}
	return customerInfo, nil
}
