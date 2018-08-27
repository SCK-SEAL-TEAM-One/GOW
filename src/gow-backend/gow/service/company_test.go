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
