package model_test

import (
	. "gow/model"
	"testing"
)

func Test_ToCustomerModel_With_Customer_Name_Th_Should_Be_Customer_Name_Th_Not_Empty(t *testing.T) {
	newCustomer := NewCustomer{
		Company: "บริษัท ที.เอ็น. อินคอร์ปอเรชั่น จำกัด",
		Branch:  "สำนักงานใหญ่",
		Address: "3 อาคารรัจนาการ ถนนสาทรใต้ แขวงยานนาวา เขตสาทร กรุงเทพมหานคร 10120",
		TaxID:   "0105553108372",
	}
	expectedCustomer := Customer{
		NameTH:    "บริษัท ที.เอ็น. อินคอร์ปอเรชั่น จำกัด",
		BranchTH:  "สำนักงานใหญ่",
		AddressTH: "3 อาคารรัจนาการ ถนนสาทรใต้ แขวงยานนาวา เขตสาทร กรุงเทพมหานคร 10120",
		TaxID:     "0105553108372",
	}

	actualCustomer := newCustomer.ToCustomerModel()

	if expectedCustomer != actualCustomer {
		t.Errorf("expect %v but it got %v", expectedCustomer, actualCustomer)
	}
}

func Test_ToCustomerModel_With_Customer_Name_En_Should_Be_Customer_Name_En_Not_Empty(t *testing.T) {
	newCustomer := NewCustomer{
		Company: "Zenity Group Co., Ltd.",
		Branch:  "Head office",
		Address: "85/28 Lat Phrao 94, Phlabphla, Want thong Lang, Bangkok",
		TaxID:   "0105560015309",
	}
	expectedCustomer := Customer{
		NameEN:    "Zenity Group Co., Ltd.",
		BranchEN:  "Head office",
		AddressEN: "85/28 Lat Phrao 94, Phlabphla, Want thong Lang, Bangkok",
		TaxID:     "0105560015309",
	}

	actualCustomer := newCustomer.ToCustomerModel()

	if expectedCustomer != actualCustomer {
		t.Errorf("expect %v but it got %v", expectedCustomer, actualCustomer)
	}
}
