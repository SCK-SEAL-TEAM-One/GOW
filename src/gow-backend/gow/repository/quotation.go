package repository

import "gow/model"

type QuotationRepository interface {
	InsertQuotation(model.QuotationForm, model.Payment, float64) (int64, error)
	GetQuotationByID(int64) (model.Quotation, error)
}
