package service_test

import "gow/model"

type MockCustomerRepository struct {
}

type MockCompanyRepository struct {
}

func (mcom MockCompanyRepository) Insert(model.NewCompany) (bool, error) { return true, nil }
func (mcom MockCompanyRepository) GetByTaxID(string) (model.CompanyInfo, error) {
	return model.CompanyInfo{
		ID:          1,
		Company:     "บริษัท สยามชำนาญกิจ จำกัด",
		Branch:      "สำนักงานใหญ่",
		Address:     "เลขที่ 3 อาคารพร้อมพันธุ์ 3 ชั้น 10 ห้อง 1001 ซอยลาดพร้าว 3 ถนนลาดพร้าว แขวงจอมพล เขตจตุจักร กรุงเทพมหานคร",
		TaxID:       "0105556042151",
		PhoneNumber: "+66979515936",
	}, nil
}
func (mcom MockCompanyRepository) GetAll() ([]model.CompanyInfo, error) {
	return []model.CompanyInfo{
		{
			ID:          1,
			Company:     "บริษัท สยามชำนาญกิจ จำกัด",
			Branch:      "สำนักงานใหญ่",
			Address:     "เลขที่ 3 อาคารพร้อมพันธ์ุ 3 ชั้น 10 ห้อง 1001 ซอยลาดพร้าม 3 ถนนลาดพร้าว แขวงจอมพล เขตจตุจักร กรุงเทพมหานคร 10900",
			TaxID:       "0705556042131",
			PhoneNumber: "+66979575936",
		},
	}, nil
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
			{
				ID:      1,
				Company: "บริษัท ที.เอ็น. อินคอร์ปอเรชั่นจำกัด",
				Branch:  "สำนักงานใหญ่",
				Address: "3 อาคารรัจนาการ ถนนสาทรใต้ แขวงยานนาวา เขตสาทร กรุงเทพมหานคร 10120",
				TaxID:   "0105553108372",
			},
		},
		nil
}
