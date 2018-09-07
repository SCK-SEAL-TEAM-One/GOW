package repository

import (
	"gow/model"
)

type QuotationRepository interface {
	InsertQuotation(quotataionForm model.QuotationForm) (int, error)
}
