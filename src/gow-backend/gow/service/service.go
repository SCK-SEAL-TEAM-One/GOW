package service

import (
	"gow/model"
)

type CustomerService interface {
	CreateNewCustomer(model.NewCustomer) (model.CustomerInfo, error)
	ListCustomers() ([]model.CustomerInfo, error)
}

type CompaniesService interface {
	CreateNewCompany(model.NewCompany) (model.CompanyInfo, error)
	GetCompanies() ([]model.CompanyInfo, error)
}
