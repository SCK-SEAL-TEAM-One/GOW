package api_test

import (
	"encoding/json"
	"gow/model"
	"io/ioutil"
)

type mockCompanyService struct {
}

type mockCustomerService struct {
}

func (mhs mockCustomerService) ListCustomers() ([]model.CustomerInfo, error) {
	var customers = []model.CustomerInfo{
		model.CustomerInfo{
			ID:      1,
			Company: "บริษัท ที.เอ็น. อินคอร์ปอเรชั่น จำกัด",
			Branch:  "สำนักงานใหญ่",
			Address: "3 อาคารรัจนาการ ถนนสาทรใต้ แขวงยานนาวา เขตสาทร กรุงเทพมหานคร 10120",
			TaxID:   "0105553108372",
		},
	}

	return customers, nil
}
func (mhs mockCustomerService) CreateNewCustomer(newcustomer model.NewCustomer) (model.CustomerInfo, error) {
	var customerInfo model.CustomerInfo
	customers, _ := ioutil.ReadFile("./customerResponse.json")
	json.Unmarshal(customers, &customerInfo)
	return customerInfo, nil
}

func (m mockCompanyService) CreateNewCompany(newCompany model.NewCompany) (model.CompanyInfo, error) {
	var companyInfo model.CompanyInfo
	company, _ := ioutil.ReadFile("./companyResponse.json")
	json.Unmarshal(company, &companyInfo)
	return companyInfo, nil
}
func (m mockCompanyService) GetCompanies() ([]model.CompanyInfo, error) {
	return []model.CompanyInfo{}, nil
}