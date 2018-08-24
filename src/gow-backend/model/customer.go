package model

import (
	"time"
)

type NewCustomer struct {
	Company string `json:"company"`
	Branch  string `json:"branch"`
	Address string `json:"address"`
	TaxID   string `json:"taxid"`
}
type CustomerInfo struct {
	ID          int       `json:"id"`
	Company     string    `json:"company"`
	Branch      string    `json:"branch"`
	Address     string    `json:"address"`
	TaxID       string    `json:"taxid"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}
