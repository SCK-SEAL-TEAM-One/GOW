package repository

import (
	"database/sql"
	"gow/model"
	"time"
)

type OrderRepository interface {
	InsertOrder(model.QuotationForm, int64) (bool, error)
	GetByQuotationID(string) ([]model.Order, error)
}

type OrderRepositoryMySQL struct {
	DBConnection *sql.DB
}

func (orderRepository OrderRepositoryMySQL) InsertOrder(quotationForm model.QuotationForm, quotationID int64) (bool, error) {
	if len(quotationForm.Orders) == 0 {
		return true, nil
	}
	statement := `INSERT INTO order 
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

func (orderRepository OrderRepositoryMySQL) GetByQuotationID(string) ([]model.Order, error) {
	return []model.Order{}, nil
}
