package model

import "time"

type NewCompany struct {
	Company     string `json:"company"`
	Branch      string `json:"branch"`
	Address     string `json:"address"`
	TaxID       string `json:"taxid"`
	PhoneNumber string `json:"phoneNumber"`
}
type CompanyInfo struct {
	ID          int       `json:"id"`
	Company     string    `json:"company"`
	Branch      string    `json:"branch"`
	Address     string    `json:"address"`
	TaxID       string    `json:"taxid"`
	PhoneNumber string    `json:"phoneNumber"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}
