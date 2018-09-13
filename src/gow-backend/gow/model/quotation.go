package model

import (
	"time"
)

type QuotationForm struct {
	CompanyID   int     `json:"companyId"`
	CustomerID  int     `json:"customerId"`
	Contact     Contact `json:"contact"`
	ProjectName string  `json:"projectName"`
	Orders      []Order `json:"orders"`
	IncludeVAT  bool    `json:"includeVAT"`
	Discount    string  `json:"payment>discount`
}

type Contact struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
type Order struct {
	OrderID      int    `json:"orderId"`
	OrderCourse  string `json:"orderCourse"`
	Amount       int    `json:"amount"`
	PricePerUnit string `json:"pricePerUnit"`
	Price        string `json:"price"`
}

type QuotationInfo struct {
	Company     CompanyQuotationInfo  `json:"company"`
	Customer    CustomerQuotationInfo `json:"customer"`
	Contact     Contact               `json:"contact"`
	ProjectName string                `json:"projectName"`
	Orders      []Order               `json:"orders"`
	Payment     Payment               `json:"payment"`
	IncludeVAT  bool                  `json:"includeVAT"`
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

type Payment struct {
	TotalPrice         string `json:"totalPrice"`
	Discount           string `json:"discount"`
	PriceAfterDiscount string `json:"priceAfterDiscount"`
	VAT                string `json:"vat"`
	NetTotalPrice      string `json:"netTotalPrice"`
	TotalPriceThai     string `json:"totalPriceThai"`
}

type Quotation struct {
	QuotationID        int
	QuotationNumber    string
	CustomerID         int
	CompanyID          int
	ContactName        string
	ContactEmail       string
	ContactPhoneNumber string
	TotalPrice         string
	Discount           string
	PriceAfterDiscount string
	VAT                string
	NetTotalPrice      string
	TotalPriceThai     string
	ProjectName        string
	QuotationDate      time.Time
	VatRate            float64
	VatIncluded        bool
}
