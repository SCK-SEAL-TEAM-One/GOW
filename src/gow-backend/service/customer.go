package service

import (
	"fmt"
	"gow-backend/model"
	"gow-backend/repository"
)

type CustomerServiceMySQL struct {
	CustomerRepository repository.CustomerRepository
}

func (customerService CustomerServiceMySQL) CreateNewCustomer(newCustomer model.NewCustomer) (model.CustomerInfo, error) {

	isCreated, err := customerService.CustomerRepository.Insert(newCustomer)
	if err != nil {
		return model.CustomerInfo{}, err
	}
	if !isCreated {
		return model.CustomerInfo{}, fmt.Errorf("can not Create")
	}

	createdCustomer, err := customerService.CustomerRepository.GetByTaxID(newCustomer.TaxID)
	if err != nil {
		return createdCustomer, err
	}
	return createdCustomer, nil

}
