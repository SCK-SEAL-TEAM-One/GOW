package service

import (
	"fmt"
	"gow/model"
	"gow/repository"
	"time"
)

const vat = 7.00
const buddistEra = 543

type QuotationServiceMySQL struct {
	QuotationRepository repository.QuotationRepository
	OrderRepository     repository.OrderRepository
	CompanyService      CompanyService
	CustomerService     CustomerService
	GetCurrentTime      func() time.Time
}

func NewQuotationServiceMySQL(quotationRepository repository.QuotationRepository, orderRepository repository.OrderRepository, companyService CompanyService, customerService CustomerService) QuotationServiceMySQL {
	return QuotationServiceMySQL{
		QuotationRepository: quotationRepository,
		OrderRepository:     orderRepository,
		CompanyService:      companyService,
		CustomerService:     customerService,
		GetCurrentTime: func() time.Time {
			return time.Now()
		},
	}
}

func (quotationService QuotationServiceMySQL) CreateQuotation(quotationForm model.QuotationForm) (int64, error) {
	// # Step 1 get quotation number
	// # Step 2 create quotation
	// # Step 3 create orders
	// # Step 4 get quotation detail
	quotationNumber := quotationService.GenerateQuotationNumber()

	quotationID, err := quotationService.QuotationRepository.InsertQuotation(quotationForm, quotationNumber, vat)
	if err != nil {
		return -1, err
	}

	_, err = quotationService.OrderRepository.InsertOrder(quotationForm, quotationID)
	if err != nil {
		return -1, err
	}

	return quotationID, nil
}

func (quotationService QuotationServiceMySQL) GenerateQuotationNumber() string {
	newQuotationNumber, _ := quotationService.QuotationRepository.GetNewQuotationNumber()
	currentTime := quotationService.GetCurrentTime()
	yearInBuddistEra := currentTime.Year() + buddistEra
	month := currentTime.Month()
	quationNumber := fmt.Sprintf("QT%d%02d-101%03d", yearInBuddistEra, month, newQuotationNumber)
	return quationNumber
}
