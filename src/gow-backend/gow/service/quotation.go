package service

import (
	"gow/model"
	"gow/repository"
)

type QuotationServiceMySQL struct {
	QuotationRepository repository.QuotationRepository
}

func (quotationService QuotationServiceMySQL) CreateQuotation(model.QuotationForm) (model.QuotationInfo, error) {
	return model.QuotationInfo{}, nil
}
