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
	customer := newCustomer.ToCustomerModel()
	statement := `INSERT INTO customer 
	(customer_name_th,customer_name_en,customer_branch_th,customer_branch_en,
	customer_address_th,customer_address_en,customer_taxid,created_time,updated_time)
	VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ? )`
	statementInsert, err := repository.DBConnection.Prepare(statement)
	if err != nil {
		return false, err
	}
	defer statementInsert.Close()
	_, err = statementInsert.Exec(customer.NameTH, customer.NameEN, customer.BranchTH, customer.BranchEN, customer.AddressTH, customer.AddressEN, customer.TaxID, time.Now(), time.Now())
	if err != nil {
		return false, err
	}
	return true, nil
}
func (repository CustomerRepositoryMySQL) GetByTaxID(customerTaxID string) (model.CustomerInfo, error) {
	var customer model.Customer
	statementQuery := `SELECT * FROM customer WHERE customer_taxid=?`
	row := repository.DBConnection.QueryRow(statementQuery, customerTaxID)
	err := row.Scan(
		&customer.ID,
		&customer.NameTH,
		&customer.NameEN,
		&customer.BranchTH,
		&customer.BranchEN,
		&customer.AddressTH,
		&customer.AddressEN,
		&customer.TaxID,
		&customer.CreatedTime,
		&customer.UpdatedTime,
	)
	if err != nil {
		return model.CustomerInfo{}, err
	}
	return customer.ToCustomerInfo(), nil
}
func (repository CustomerRepositoryMySQL) GetAll() ([]model.CustomerInfo, error) {
	var customerInfo []model.CustomerInfo
	statementQuery := `SELECT * FROM customer `
	rows, err := repository.DBConnection.Query(statementQuery)
	if err != nil {
		return customerInfo, err
	}
	for rows.Next() {
		var customer model.Customer
		rows.Scan(
			&customer.ID,
			&customer.NameTH,
			&customer.NameEN,
			&customer.BranchTH,
			&customer.BranchEN,
			&customer.AddressTH,
			&customer.AddressEN,
			&customer.TaxID,
			&customer.CreatedTime,
			&customer.UpdatedTime,
		)
		customerInfo = append(customerInfo, customer.ToCustomerInfo())
	}
	return customerInfo, nil
}
