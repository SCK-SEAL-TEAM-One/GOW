package service

import (
	"gow/model"
)

type CustomerService interface {
	CreateNewCustomer(model.NewCustomer) (model.CustomerInfo, error)
	ListCustomers() ([]model.CustomerInfo, error)
	GetCustomerByTaxID(string) (model.CustomerInfo, error)
}

type CompanyService interface {
	CreateNewCompany(model.NewCompany) (model.CompanyInfo, error)
	ListCompanies() ([]model.CompanyInfo, error)
	GetCompanyByTaxID(string) (model.CompanyInfo, error)
}

type QuotationService interface {
	CreateQuotation(model.QuotationForm) (int64, error)
	GetQuotationByQuotationNumber(string) (model.QuotationInfo, error)
	GetQuotationByQuotationID(int64) (model.QuotationInfo, error)
}
