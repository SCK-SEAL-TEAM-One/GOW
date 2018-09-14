package service_test

import (
	"gow/model"
	. "gow/service"
	"testing"
)

func Test_CreateNewCustomer_Input_TN_Corp_Should_Be_Created_Customer_Name_TN_Corp(t *testing.T) {
	customerService := CustomerServiceMySQL{
		CustomerRepository: &MockCustomerRepository{},
	}
	expectedCustomer := model.CustomerInfo{
		ID:      1,
		Company: "บริษัท ที.เอ็น. อินคอร์ปอเรชั่นจำกัด",
		Branch:  "สำนักงานใหญ่",
		Address: "3 อาคารรัจนาการ ถนนสาทรใต้ แขวงยานนาวา เขตสาทร กรุงเทพมหานคร 10120",
		TaxID:   "0105553108372",
	}
	newCustomer := model.NewCustomer{
		Company: "บริษัท ที.เอ็น. อินคอร์ปอเรชั่นจำกัด",
		Branch:  "สำนักงานใหญ่",
		Address: "3 อาคารรัจนาการ ถนนสาทรใต้ แขวงยานนาวา เขตสาทร กรุงเทพมหานคร 10120",
		TaxID:   "0105553108372",
	}

	createdCustomer, _ := customerService.CreateNewCustomer(newCustomer)

	if expectedCustomer != createdCustomer {
		t.Errorf("expect %v but got it %v", expectedCustomer, createdCustomer)
	}
}

func Test_ListCustomers_Should_Be_CustomerInfo(t *testing.T) {
	customerService := CustomerServiceMySQL{
		CustomerRepository: &MockCustomerRepository{},
	}
	expectedCustomers := []model.CustomerInfo{
		{ID: 1,
			Company: "บริษัท ที.เอ็น. อินคอร์ปอเรชั่นจำกัด",
			Branch:  "สำนักงานใหญ่",
			Address: "3 อาคารรัจนาการ ถนนสาทรใต้ แขวงยานนาวา เขตสาทร กรุงเทพมหานคร 10120",
			TaxID:   "0105553108372",
		},
	}

	listCustomers, _ := customerService.ListCustomers()
	for index, listCustomer := range listCustomers {
		if expectedCustomers[index] != listCustomer {
			t.Errorf("expect %v but got it %v", expectedCustomers, listCustomers)
		}
	}

}

func Test_GetCustomerByTaxID_Input_ID_0105537133562_Should_Be_CompanyInfo_Aycap(t *testing.T) {
	customerService := CustomerServiceMySQL{
		CustomerRepository: &MockCustomerRepository{},
	}
	customerTaxId := "0105537133562"
	expectedCustomer := model.CustomerInfo{
		ID:      1,
		Company: "บริษัท อยุธยา แคปปิตอล เซอร์วิสเซ",
		Branch:  "สำนักงานใหญ่",
		Address: "อาคารกรุงศรีเพลินจิต ทาวเวอร์ 550 ถนนเพลินจิต แขวงเขตปทุมวัน กรุงเทพมหานคร",
		TaxID:   "0105537133562",
	}

	actualCustomer, _ := customerService.GetCustomerByTaxID(customerTaxId)

	if expectedCustomer != actualCustomer {
		t.Errorf("expect %v but got it %v", expectedCustomer, actualCustomer)
	}

}
