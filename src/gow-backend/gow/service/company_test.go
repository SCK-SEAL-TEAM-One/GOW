package service_test

import (
	"gow/model"
	. "gow/service"
	"testing"
)

func Test_CreateNewCompany_Input_Data_Companies_SiamChamnankit_Should_Be_CompaniesInfo_Error(t *testing.T) {
	companyService := CompanyServiceMySQL{
		CompanyRepository: &MockCompanyRepository{},
	}
	expectedCompany := model.CompanyInfo{
		ID:          1,
		Company:     "บริษัท สยามชำนาญกิจ จำกัด",
		Branch:      "สำนักงานใหญ่",
		Address:     "เลขที่ 3 อาคารพร้อมพันธุ์ 3 ชั้น 10 ห้อง 1001 ซอยลาดพร้าว 3 ถนนลาดพร้าว แขวงจอมพล เขตจตุจักร กรุงเทพมหานคร",
		TaxID:       "0105556042151",
		PhoneNumber: "+66979515936",
	}
	newCompany := model.NewCompany{
		Company:     "บริษัท สยามชำนาญกิจ จำกัด",
		Branch:      "สำนักงานใหญ่",
		Address:     "เลขที่ 3 อาคารพร้อมพันธุ์ 3 ชั้น 10 ห้อง 1001 ซอยลาดพร้าว 3 ถนนลาดพร้าว แขวงจอมพล เขตจตุจักร กรุงเทพมหานคร",
		TaxID:       "0105556042151",
		PhoneNumber: "+66979515936",
	}

	createdCompany, _ := companyService.CreateNewCompany(newCompany)

	if expectedCompany != createdCompany {
		t.Errorf("expect %v but got it %v", expectedCompany, createdCompany)
	}

}

func Test_ListCompanies_Should_Be_CompaniesInfo(t *testing.T) {
	companyService := CompanyServiceMySQL{
		CompanyRepository: &MockCompanyRepository{},
	}
	expectedCompanies := []model.CompanyInfo{
		{
			ID:          1,
			Company:     "บริษัท สยามชำนาญกิจ จำกัด",
			Branch:      "สำนักงานใหญ่",
			Address:     "เลขที่ 3 อาคารพร้อมพันธ์ุ 3 ชั้น 10 ห้อง 1001 ซอยลาดพร้าม 3 ถนนลาดพร้าว แขวงจอมพล เขตจตุจักร กรุงเทพมหานคร 10900",
			TaxID:       "0705556042131",
			PhoneNumber: "+66979575936",
		},
	}

	listCompanies, _ := companyService.ListCompanies()

	if len(listCompanies) == 0 {
		t.Errorf("expected companies length is %d but it got %d", len(expectedCompanies), len(listCompanies))
	}
	for index, listCompany := range listCompanies {
		if expectedCompanies[index] != listCompany {
			t.Errorf("expect %v but got it %v", expectedCompanies, listCompanies)
		}
	}
}

func Test_ConvertMoneyToThaiCharactor_Input_100000_Should_Be_One_Hundred_Thousand_Baht(t *testing.T) {
	expectedNumber := "หนึ่งแสนบาทถ้วน"
	number := 100000.00

	actual := ConvertMoneyToThaiCharactor(number)

	if expectedNumber != actual {
		t.Errorf("expect %s but it got %s", expectedNumber, actual)
	}
}

func Test_GetCompanyByTaxID_Input_ID_2_Should_Be_CompanyInfo_Aycap(t *testing.T) {
	companyService := CompanyServiceMySQL{
		CompanyRepository: &MockCompanyRepository{},
	}
	companyTaxId := "0105561001221"
	expectedCompany := model.CompanyInfo{
		ID:          2,
		Company:     "บริษัท ชู ฮา ริ จำกัด",
		Branch:      "สำนักงานใหญ่",
		Address:     "เลขที่ 3 อาคารพร้อมพันธ์ุ 3 ชั้น 10 ห้อง 1002 ซอยลาดพร้าม 3 ถนนลาดพร้าว แขวงจอมพล เขตจตุจักร กรุงเทพมหานคร 10900",
		TaxID:       "0105561001221",
		PhoneNumber: "+66979575936",
	}

	actualCompany, _ := companyService.GetCompanyByTaxID(companyTaxId)

	if expectedCompany != actualCompany {
		t.Errorf("expect %v but got it %v", expectedCompany, actualCompany)
	}

}
