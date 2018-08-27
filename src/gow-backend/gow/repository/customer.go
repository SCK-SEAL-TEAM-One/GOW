package repository

import (
	"database/sql"
	"gow/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepository interface {
	Insert(model.NewCustomer) (bool, error)
	GetByTaxID(string) (model.CustomerInfo, error)
	GetAll() ([]model.CustomerInfo, error)
}
type CustomerRepositoryMySQL struct {
	DBConnection *sql.DB
}

func (repository CustomerRepositoryMySQL) Insert(newCustomer model.NewCustomer) (bool, error) {
	statementInsert, err := repository.DBConnection.Prepare(`INSERT INTO customer 
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
	row := repository.DBConnection.QueryRow(statementQuery, customerTaxID)
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
func (repository CustomerRepositoryMySQL) GetAll() ([]model.CustomerInfo, error) {
	var customerInfo []model.CustomerInfo
	statementQuery := `SELECT * FROM customer `
	rows, err := repository.DBConnection.Query(statementQuery)
	if err != nil {
		return customerInfo, err
	}
	for rows.Next() {
		var customer model.CustomerInfo
		rows.Scan(
			&customer.ID,
			&customer.Company,
			&customer.Branch,
			&customer.Address,
			&customer.TaxID,
			&customer.CreatedTime,
			&customer.UpdatedTime,
		)
		customerInfo = append(customerInfo, customer)
	}
	return customerInfo, nil
}
