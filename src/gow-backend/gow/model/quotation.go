package model

type QuotationForm struct {
	CompanyID   string  `json:"companyId"`
	CustomerID  string  `json:"customerId"`
	Contact     Contact `json:"contact"`
	ProjectName string  `json:"projectName"`
	orders      []Order `json:"orders"`
	IncludeVAT  bool    `json:"includeVAT"`
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
	Discount           string `json:"discount"`
	PriceAfterDiscount string `json:"priceAfterDiscount"`
	VAT                string `json:"vat"`
	TotalPrice         string `json:"totalPrice"`
	TotalPriceThai     string `json:"totalPriceThai"`
}
