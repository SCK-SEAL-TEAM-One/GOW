package service

import (
	"fmt"
	"gow/model"
	"gow/repository"
)

type CompanyServiceMySQL struct {
	CompanyRepository repository.CompanyRepository
}

func (companyService CompanyServiceMySQL) CreateNewCompany(newCompany model.NewCompany) (model.CompanyInfo, error) {
	isCreated, err := companyService.CompanyRepository.Insert(newCompany)
	if err != nil {
		return model.CompanyInfo{}, err
	}
	if !isCreated {
		return model.CompanyInfo{}, fmt.Errorf("can not Create")
	}

	createdCompany, err := companyService.CompanyRepository.GetByTaxID(newCompany.TaxID)
	if err != nil {
		return createdCompany, err
	}
	return createdCompany, nil
}

func (companyService CompanyServiceMySQL) ListCompanies() ([]model.CompanyInfo, error) {
	return companyService.CompanyRepository.GetAll()
}

func (companyService CompanyServiceMySQL) GetCompanyByID(int) (model.CompanyInfo, error) {
	return model.CompanyInfo{}, nil
}
