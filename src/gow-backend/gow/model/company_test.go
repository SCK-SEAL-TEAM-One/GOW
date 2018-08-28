package model_test

import (
	. "gow/model"
	"testing"
)

func Test_ToCompanyModel_With_Company_Name_Th_Should_Be_Company_Name_Th_Not_Empty(t *testing.T) {
	newCompany := NewCompany{
		Company:     "บริษัทสยามชำนาญกิจ จำกัด",
		Branch:      "สำนักงานใหญ่",
		Address:     "เลขที่ 3 อาคารพร้อมพันธ์ุ 3 ชั้น 10 ห้อง 1001 ซอยลาดพร้าม 3 ถนนลาดพร้าว แขวงจอมพล เขตจตุจักร กรุงเทพมหานคร 10900",
		TaxID:       "0105556042151",
		PhoneNumber: "+66979575936",
	}

	expectedCoompany := Company{
		NameTH:      "บริษัทสยามชำนาญกิจ จำกัด",
		BranchTH:    "สำนักงานใหญ่",
		AddressTH:   "เลขที่ 3 อาคารพร้อมพันธ์ุ 3 ชั้น 10 ห้อง 1001 ซอยลาดพร้าม 3 ถนนลาดพร้าว แขวงจอมพล เขตจตุจักร กรุงเทพมหานคร 10900",
		TaxID:       "0105556042151",
		PhoneNumber: "+66979575936",
	}

	actualCompany := newCompany.ToCompanyModel()

	if expectedCoompany != actualCompany {
		t.Errorf("expect %v but it got %v", expectedCoompany, actualCompany)
	}
}
func Test_ToCompanyModel_With_Company_Name_En_Should_Be_Company_Name_En_Not_Empty(t *testing.T) {
	newCompany := NewCompany{
		Company:     "Siam Chamnankit Company Limited",
		Branch:      "Head office",
		Address:     "No.3 Promphan Building 3, 10th floor,Room 1001 Soi Lat Phrao 3, Lat Phrao Road, Chom Phon, Chatuchak, Bangkok, 10900",
		TaxID:       "0105556042151",
		PhoneNumber: "+66979575936",
	}

	expectedCoompany := Company{
		NameEN:      "Siam Chamnankit Company Limited",
		BranchEN:    "Head office",
		AddressEN:   "No.3 Promphan Building 3, 10th floor,Room 1001 Soi Lat Phrao 3, Lat Phrao Road, Chom Phon, Chatuchak, Bangkok, 10900",
		TaxID:       "0105556042151",
		PhoneNumber: "+66979575936",
	}

	actualCompany := newCompany.ToCompanyModel()

	if expectedCoompany != actualCompany {
		t.Errorf("expect %v but it got %v", expectedCoompany, actualCompany)
	}
}
