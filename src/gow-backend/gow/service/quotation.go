package service

import (
	"gow/model"
	"gow/repository"
	"strconv"
	"strings"
)

type QuotationServiceMySQL struct {
	QuotationRepository repository.QuotationRepository
	OrderRepository     repository.OrderRepository
	CompanyService      CompanyService
	CustomerService     CustomerService
}

const vat = 7

func (quotationService QuotationServiceMySQL) CreateQuotation(quotationForm model.QuotationForm) (model.QuotationInfo, error) {
	totalPrice, err := CalculateOrdersPrice(&quotationForm.Orders)
	if err != nil {
		return model.QuotationInfo{}, err
	}

	discount := strings.Replace(quotationForm.Payment.Discount, ",", "", -1)
	discountFloat, err := strconv.ParseFloat(discount, 64)
	if err != nil {
		return model.QuotationInfo{}, err
	}

	priceAfterDiscount := CalculateDiscount(totalPrice, discountFloat)
	taxFee := CalculateVat(priceAfterDiscount, vat)
	netTotalPrice := CalculateNetTotalPrice(priceAfterDiscount, taxFee)

	payment := model.Payment{
		TotalPrice:         AddComma(totalPrice),
		Discount:           AddComma(discountFloat),
		PriceAfterDiscount: AddComma(priceAfterDiscount),
		VAT:                AddComma(taxFee),
		NetTotalPrice:      AddComma(netTotalPrice),
		TotalPriceThai:     ConvertMoneyToThaiCharactor(netTotalPrice),
	}

	quotationID, err := quotationService.QuotationRepository.InsertQuotation(quotationForm, payment, float64(vat))
	if err != nil {
		return model.QuotationInfo{}, err
	}

	_, err = quotationService.OrderRepository.InsertOrder(quotationForm, quotationID)
	if err != nil {
		return model.QuotationInfo{}, err
	}

	quotation, err := quotationService.QuotationRepository.GetQuotationByID(quotationID)
	if err != nil {
		return model.QuotationInfo{}, err
	}

	company, err := quotationService.CompanyService.GetCompanyByID(quotation.CompanyID)
	if err != nil {
		return model.QuotationInfo{}, err
	}

	customer, err := quotationService.CustomerService.GetCustomerByID(quotation.CustomerID)
	if err != nil {
		return model.QuotationInfo{}, err
	}

	return model.QuotationInfo{
		Company:     company.ToCompanyQuotationInfo(),
		Customer:    customer.ToCustomerQuotationInfo(),
		Orders:      quotationForm.Orders,
		Payment:     payment,
		Contact:     quotationForm.Contact,
		ProjectName: quotationForm.ProjectName,
		IncludeVAT:  quotationForm.IncludeVAT,
	}, nil
}
