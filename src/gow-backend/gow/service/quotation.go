package service

import "gow/model"

type QuotationService interface {
	CreateQuotation(model.QuotationForm) (model.QuotationInfo, error)
}
