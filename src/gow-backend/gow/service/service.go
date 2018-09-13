package service

import (
	"gow/model"
)

type CustomerService interface {
	CreateNewCustomer(model.NewCustomer) (model.CustomerInfo, error)
	ListCustomers() ([]model.CustomerInfo, error)
	GetCustomerByID(int) (model.CustomerInfo, error)
}

type CompanyService interface {
	CreateNewCompany(model.NewCompany) (model.CompanyInfo, error)
	ListCompanies() ([]model.CompanyInfo, error)
	GetCompanyByID(int) (model.CompanyInfo, error)
}

type QuotationService interface {
	CreateQuotation(model.QuotationForm) (model.QuotationInfo, error)
}
