package model

import (
	"time"
)

type NewCustomer struct {
	Company string
	Branch  string
	Address string
	TaxID   string
}
type CustomerInfo struct {
	ID          int       `json:"id"`
	Company     string    `json:"company"`
	Branch      string    `json:"beanch"`
	Address     string    `json:"address"`
	TaxID       string    `json:"taxID"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}
