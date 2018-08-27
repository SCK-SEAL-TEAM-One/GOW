package service_test

import "gow-backend/model"

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
