package repository

import "gow-backend/model"

type CustomerRepository interface {
	Insert(model.NewCustomer) (bool, error)
	GetByTaxID(string) (model.CustomerInfo, error)
}
type CustomerRepositoryMySQL struct {
}

func (m CustomerRepositoryMySQL) Insert(model.NewCustomer) (bool, error) { return true, nil }
func (m CustomerRepositoryMySQL) GetByTaxID(string) (model.CustomerInfo, error) {
	return model.CustomerInfo{
		ID:      1,
		Company: "บริษัท ที.เอ็น. อินคอร์ปอเรชั่นจำกัด",
		Branch:  "สำนักงานใหญ่",
		Address: "3 อาคารรัจนาการ ถนนสาทรใต้ แขวงยานนาวา เขตสาทร กรุงเทพมหานคร 10120",
		TaxID:   "0105553108372",
	}, nil
}
