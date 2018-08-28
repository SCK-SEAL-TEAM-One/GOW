package model

import (
	"regexp"
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
type Customer struct {
	ID          int
	NameTH      string
	NameEN      string
	BranchTH    string
	BranchEN    string
	AddressTH   string
	AddressEN   string
	TaxID       string
	CreatedTime time.Time
	UpdatedTime time.Time
}

func (newCustomer NewCustomer) ToCustomerModel() Customer {
	isThaiWord, _ := regexp.MatchString("[ก-ฮ]+", newCustomer.Company)
	if isThaiWord {
		return Customer{
			NameTH:    newCustomer.Company,
			BranchTH:  newCustomer.Branch,
			AddressTH: newCustomer.Address,
			TaxID:     newCustomer.TaxID,
		}
	}
	return Customer{
		NameEN:    newCustomer.Company,
		BranchEN:  newCustomer.Branch,
		AddressEN: newCustomer.Address,
		TaxID:     newCustomer.TaxID,
	}
}

func (customer Customer) ToCustomerInfo() CustomerInfo {
	if customer.NameTH != "" {
		return CustomerInfo{
			ID:          customer.ID,
			Company:     customer.NameTH,
			Branch:      customer.BranchTH,
			Address:     customer.AddressTH,
			TaxID:       customer.TaxID,
			CreatedTime: customer.CreatedTime,
			UpdatedTime: customer.UpdatedTime,
		}
	}
	return CustomerInfo{
		ID:          customer.ID,
		Company:     customer.NameEN,
		Branch:      customer.BranchEN,
		Address:     customer.AddressEN,
		TaxID:       customer.TaxID,
		CreatedTime: customer.CreatedTime,
		UpdatedTime: customer.UpdatedTime,
	}
}
