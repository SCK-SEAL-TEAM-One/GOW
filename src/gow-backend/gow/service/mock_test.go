package service_test

import (
	"encoding/json"
	"gow/model"
	"io/ioutil"
)

type MockCustomerRepository struct {
}

type MockCompanyRepository struct {
}

func (mcom MockCompanyRepository) Insert(model.NewCompany) (bool, error) { return true, nil }
func (mcom MockCompanyRepository) GetByTaxID(companyTaxID string) (model.CompanyInfo, error) {
	if companyTaxID == "0105561001221" {
		return model.CompanyInfo{
			ID:          2,
			Company:     "บริษัท ชู ฮา ริ จำกัด",
			Branch:      "สำนักงานใหญ่",
			Address:     "เลขที่ 3 อาคารพร้อมพันธ์ุ 3 ชั้น 10 ห้อง 1002 ซอยลาดพร้าม 3 ถนนลาดพร้าว แขวงจอมพล เขตจตุจักร กรุงเทพมหานคร 10900",
			TaxID:       "0105561001221",
			PhoneNumber: "+66979575936",
		}, nil
	}
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
func (m MockCustomerRepository) GetByTaxID(taxID string) (model.CustomerInfo, error) {
	if taxID == "0105553108372" {
		return model.CustomerInfo{
			ID:      1,
			Company: "บริษัท ที.เอ็น. อินคอร์ปอเรชั่นจำกัด",
			Branch:  "สำนักงานใหญ่",
			Address: "3 อาคารรัจนาการ ถนนสาทรใต้ แขวงยานนาวา เขตสาทร กรุงเทพมหานคร 10120",
			TaxID:   "0105553108372",
		}, nil
	}
	return model.CustomerInfo{
		ID:      1,
		Company: "บริษัท อยุธยา แคปปิตอล เซอร์วิสเซ",
		Branch:  "สำนักงานใหญ่",
		Address: "อาคารกรุงศรีเพลินจิต ทาวเวอร์ 550 ถนนเพลินจิต แขวงเขตปทุมวัน กรุงเทพมหานคร",
		TaxID:   "0105537133562",
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

type mockQuotationRepository struct {
	QuotationNumber string
	QuotationForm   model.QuotationForm
	Vat             float64
}

func (mock *mockQuotationRepository) InsertQuotation(quotationForm model.QuotationForm, quotationNumber string, vat float64) (int64, error) {
	mock.QuotationNumber = quotationNumber
	mock.QuotationForm = quotationForm
	mock.Vat = vat
	return int64(1), nil
}

func (mock mockQuotationRepository) GetQuotationByID(int64) (model.Quotation, error) {
	return model.Quotation{
		QuotationID:        1,
		QuotationNumber:    "QT256104-000002",
		ContactName:        "Nopparat Slisatkorn",
		ContactEmail:       "nopparat.slisatkorn@krungsri.com",
		ContactPhoneNumber: "",
		ProjectName:        "โครงการ 3 Days Software Testing in Action Workshop",
		VatIncluded:        false,
	}, nil
}
func (mock mockQuotationRepository) GetByQuotationNumber(string) (model.Quotation, error) {
	return model.Quotation{
		QuotationID:        1,
		QuotationNumber:    "QT256104-101002",
		TotalPrice:         100000.00,
		Discount:           0.00,
		PriceAfterDiscount: 100000.00,
		VAT:                7000.00,
		NetTotalPrice:      107000.00,
		TotalPriceThai:     "หนึ่งแสนเจ็ดพันบาทถ้วน",
		ContactName:        "Nopparat Slisatkorn",
		ContactEmail:       "nopparat.slisatkorn@krungsri.com",
		ContactPhoneNumber: "",
		ProjectName:        "โครงการ 3 Days Software Testing in Action Workshop",
		VatIncluded:        false,
	}, nil
}

func (mock mockQuotationRepository) GetNewQuotationNumber() (int, error) {
	return 2, nil
}

type mockOrderRepository struct{}

func (mock mockOrderRepository) InsertOrder(model.QuotationForm, int64) (bool, error) {
	return true, nil
}

func (mock mockOrderRepository) GetByQuotationID(int) ([]model.Order, error) {
	return []model.Order{
		{
			OrderCourse:  "ค่าฝึกอบรม Software Testing in Action จำนวน 3 วัน วันพุธที่ 20 - วันศุกร์ที่ 22 มิถุนายน พ.ศ.2561",
			Amount:       1,
			PricePerUnit: 100000.00,
			Price:        100000.00,
		},
	}, nil
}

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

func (m mockCustomerService) GetCustomerByTaxID(string) (model.CustomerInfo, error) {
	return model.CustomerInfo{
		ID:      3,
		Company: "บริษัท อยุธยา แคปปิตอล เซอร์วิสเซส จำกัด",
		Branch:  "สำนักงานใหญ่",
		Address: "อาคารกรุงศรีเพลินจิต ทาวเวอร์ 550 ถนนเพลินจิต แขวงเขตปทุมวัน กรุงเทพมหานคร",
		TaxID:   "0105537133562",
	}, nil
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

func (m mockCompanyService) GetCompanyByTaxID(string) (model.CompanyInfo, error) {
	return model.CompanyInfo{
		ID:          1,
		Company:     "บริษัท สยามชำนาญกิจ จำกัด",
		Branch:      "สำนักงานใหญ่",
		Address:     "เลขที่ 3 อาคารพร้อมพันธุ์ 3 ชั้น 10 ห้อง 1001 ซอยลาดพร้าว 3 ถนนลาดพร้าว แขวงจอมพล เขตจตุจักร กรุงเทพมหานคร",
		TaxID:       "0105556042151",
		PhoneNumber: "+66979515936",
	}, nil
}
