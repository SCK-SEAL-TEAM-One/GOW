package model

import (
	"gow/stringutil"
	"time"
)

const (
	InterestRate = 100.00
)

type Contact struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
type Order struct {
	OrderID      int     `json:"orderId"`
	OrderCourse  string  `json:"orderCourse"`
	Amount       int     `json:"amount"`
	PricePerUnit float64 `json:"pricePerUnit"`
	Price        float64 `json:"price"`
	createdTime  time.Time
	updatedTime  time.Time
}

func (order Order) GetPrice() float64 {
	return float64(order.Amount) * order.PricePerUnit
}

type Payment struct {
	TotalPrice         float64 `json:"totalPrice"`
	Discount           float64 `json:"discount"`
	PriceAfterDiscount float64 `json:"priceAfterDiscount"`
	VAT                float64 `json:"vat"`
	NetTotalPrice      float64 `json:"netTotalPrice"`
	TotalPriceThai     string  `json:"totalPriceThai"`
}
type QuotationForm struct {
	CompanyTaxID  string  `json:"companyTaxId"`
	CustomerTaxID string  `json:"customerTaxId"`
	Contact       Contact `json:"contact"`
	ProjectName   string  `json:"projectName"`
	Orders        []Order `json:"orders"`
	IncludeVAT    bool    `json:"includeVAT"`
	Payment       Payment `json:"payment"`
	VATRate       float64 `json:"vatRate"`
}

func (quotationForm QuotationForm) GetTotalPrice() float64 {
	var totalPrice float64
	for _, order := range quotationForm.Orders {
		totalPrice += order.GetPrice()
	}
	return totalPrice
}

func (quotationForm QuotationForm) GetPriceAfterDiscount() float64 {
	return quotationForm.GetTotalPrice() - quotationForm.Payment.Discount
}

func (quotationForm QuotationForm) GetVATFee() float64 {
	return quotationForm.GetPriceAfterDiscount() * quotationForm.VATRate / InterestRate
}

func (quotationForm QuotationForm) GetNetTotalPrice() float64 {
	return quotationForm.GetPriceAfterDiscount() + quotationForm.GetVATFee()
}

func (quotationForm QuotationForm) GetTotalPriceThai() string {
	return stringutil.ConvertMoneyToThaiCharactor(quotationForm.GetNetTotalPrice())
}

type CompanyQuotationInfo struct {
	ID          int    `json:"id"`
	Company     string `json:"company"`
	Branch      string `json:"branch"`
	Address     string `json:"address"`
	TaxID       string `json:"taxid"`
	PhoneNumber string `json:"phoneNumber"`
}
type CustomerQuotationInfo struct {
	ID      int    `json:"id"`
	Company string `json:"company"`
	Branch  string `json:"branch"`
	Address string `json:"address"`
	TaxID   string `json:"taxid"`
}

type QuotationInfo struct {
	Company         CompanyQuotationInfo  `json:"company"`
	Customer        CustomerQuotationInfo `json:"customer"`
	QuotationNumber string                `json:"quotationNumber"`
	Contact         Contact               `json:"contact"`
	ProjectName     string                `json:"projectName"`
	Orders          []Order               `json:"orders"`
	Payment         Payment               `json:"payment"`
	IncludeVAT      bool                  `json:"includeVAT"`
}

type Quotation struct {
	QuotationID        int
	QuotationNumber    string
	CustomerTaxID      string
	CompanyTaxID       string
	ContactName        string
	ContactEmail       string
	ContactPhoneNumber string
	TotalPrice         float64
	Discount           float64
	PriceAfterDiscount float64
	VAT                float64
	NetTotalPrice      float64
	TotalPriceThai     string
	ProjectName        string
	QuotationDate      time.Time
	VatRate            float64
	VatIncluded        bool
	CreatedTime        time.Time
	UpdatedTime        time.Time
}
