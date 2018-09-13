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
		{
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

func (m mockCustomerService) GetCustomerByID(int) (model.CustomerInfo, error) {
	return model.CustomerInfo{}, nil
}

func (m mockCompanyService) CreateNewCompany(newCompany model.NewCompany) (model.CompanyInfo, error) {
	var companyInfo model.CompanyInfo
	company, _ := ioutil.ReadFile("./companyResponse.json")
	json.Unmarshal(company, &companyInfo)
	return companyInfo, nil
}
func (m mockCompanyService) ListCompanies() ([]model.CompanyInfo, error) {
	readFlie, _ := ioutil.ReadFile("companiesResponseInfo.json")
	var companies []model.CompanyInfo
	json.Unmarshal(readFlie, &companies)
	return companies, nil
}

func (m mockCompanyService) GetCompanyByID(int) (model.CompanyInfo, error) {
	return model.CompanyInfo{}, nil
}

type mockQuotationService struct {
}

func (m mockQuotationService) CreateQuotation(model.QuotationForm) (model.QuotationInfo, error) {
	readFlie, _ := ioutil.ReadFile("quotationInfo.json")
	var quotationInfo model.QuotationInfo
	json.Unmarshal(readFlie, &quotationInfo)
	return quotationInfo, nil
}
