package service

import (
	"gow/model"
	"gow/repository"
)

type QuotationServiceMySQL struct {
	QuotationRepository repository.QuotationRepository
}

func (quotationServiceMySQL QuotationServiceMySQL) CreateQuotation(model.QuotationForm) (model.QuotationInfo, error) {
	return model.QuotationInfo{}, nil
}

func CalculatePrice(amount int, pricePerUnit float64) float64 {
	price := float64(amount) * pricePerUnit
	return price
}

func CalculateDiscount(price, discount float64) float64 {
	priceAfterDiscount := price - discount
	return priceAfterDiscount
}
