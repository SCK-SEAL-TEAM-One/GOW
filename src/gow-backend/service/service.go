package service

import (
	"gow-backend/model"
)

type CustomerService interface {
	CreateNewCustomer(model.NewCustomer) (model.CustomerInfo, error)
	GetCustomers() ([]model.CustomerInfo, error)
}
