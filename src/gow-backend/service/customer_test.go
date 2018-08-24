package service

import (
	"gow-backend/model"
	"testing"
)

type MockCustomerRepository struct {
}

func (m MockCustomerRepository) Insert(model.NewCustomer) (bool, error) { return true, nil }
func (m MockCustomerRepository) GetByTaxID(string) (model.CustomerInfo, error) {
	return model.CustomerInfo{
		ID:      1,
		Company: "บริษัท ที.เอ็น. อินคอร์ปอเรชั่นจำกัด",
		Branch:  "สำนักงานใหญ่",
		Address: "3 อาคารรัจนาการ ถนนสาทรใต้ แขวงยานนาวา เขตสาทร กรุงเทพมหานคร 10120",
		TaxID:   "0105553108372",
	}, nil
}
func (m MockCustomerRepository) GetAll() ([]model.CustomerInfo, error) {
	return []model.CustomerInfo{
			model.CustomerInfo{
				ID:      1,
				Company: "บริษัท ที.เอ็น. อินคอร์ปอเรชั่นจำกัด",
				Branch:  "สำนักงานใหญ่",
				Address: "3 อาคารรัจนาการ ถนนสาทรใต้ แขวงยานนาวา เขตสาทร กรุงเทพมหานคร 10120",
				TaxID:   "0105553108372",
			},
		},
		nil
}

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
		model.CustomerInfo{ID: 1,
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
