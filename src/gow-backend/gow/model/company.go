package model

import (
	"regexp"
	"time"
)

type NewCompany struct {
	Company     string `json:"company"`
	Branch      string `json:"branch"`
	Address     string `json:"address"`
	TaxID       string `json:"taxid"`
	PhoneNumber string `json:"phoneNumber"`
}
type Company struct {
	ID          int
	NameTH      string
	NameEN      string
	BranchTH    string
	BranchEN    string
	AddressTH   string
	AddressEN   string
	TaxID       string
	PhoneNumber string
	CreatedTime time.Time
	UpdatedTime time.Time
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

func (newCompany NewCompany) ToCompanyModel() Company {
	isThaiWord, _ := regexp.MatchString("[ก-ฮ]+", newCompany.Company)
	if isThaiWord {
		return Company{
			NameTH:      newCompany.Company,
			BranchTH:    newCompany.Branch,
			AddressTH:   newCompany.Address,
			TaxID:       newCompany.TaxID,
			PhoneNumber: newCompany.PhoneNumber,
		}
	}
	return Company{
		NameEN:      newCompany.Company,
		BranchEN:    newCompany.Branch,
		AddressEN:   newCompany.Address,
		TaxID:       newCompany.TaxID,
		PhoneNumber: newCompany.PhoneNumber,
	}
}
