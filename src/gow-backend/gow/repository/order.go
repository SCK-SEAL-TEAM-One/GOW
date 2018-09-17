package repository

import (
	"database/sql"
	"gow/model"
	"time"
)

type OrderRepository interface {
	InsertOrder(model.QuotationForm, int64) (bool, error)
	GetByQuotationID(int) ([]model.Order, error)
}

type OrderRepositoryMySQL struct {
	DBConnection *sql.DB
}

func (orderRepository OrderRepositoryMySQL) InsertOrder(quotationForm model.QuotationForm, quotationID int64) (bool, error) {
	if len(quotationForm.Orders) == 0 {
		return true, nil
	}
	statement := `INSERT INTO orders 
	(amount,price_per_unit,price,description,quotation_id, created_time, updated_time)
	VALUES ( ?, ?, ?, ?, ?, ?, ?)`
	statementInsert, err := orderRepository.DBConnection.Prepare(statement)
	if err != nil {
		return false, err
	}
	defer statementInsert.Close()
	for _, order := range quotationForm.Orders {
		_, err = statementInsert.Exec(
			order.Amount,
			order.PricePerUnit,
			order.Price,
			order.OrderCourse,
			quotationID,
			time.Now(),
			time.Now(),
		)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (orderRepository OrderRepositoryMySQL) GetByQuotationID(quotationId int) ([]model.Order, error) {
	var orders []model.Order
	statementQuery := `SELECT amount, price_per_unit, price, description FROM orders WHERE quotation_id=?`
	rows, err := orderRepository.DBConnection.Query(statementQuery, quotationId)
	if err != nil {
		return orders, err
	}

	for rows.Next() {
		var order model.Order
		rows.Scan(
			&order.Amount,
			&order.PricePerUnit,
			&order.Price,
			&order.OrderCourse,
		)
		orders = append(orders, order)
	}
	return orders, nil
}
