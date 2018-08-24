package model

type CustomerInfo struct {
	ID      int    `json:"id"`
	Company string `json:"company"`
	Branch  string `json:"branch"`
	Address string `json:"address"`
	TaxID   string `json:"taxid"`
}

type NewCustomer struct {
	Company string
	Branch  string
	Address string
	TaxID   string
}
