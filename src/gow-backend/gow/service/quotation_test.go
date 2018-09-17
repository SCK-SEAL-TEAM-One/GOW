package service_test

import (
	"gow/model"
	. "gow/service"
	"testing"
)

func Test_CreateQuotation_Input_QuotationForm_Should_Be_QuotationInfo_With_No_Error(t *testing.T) {
	expectedQuotationInfo := model.QuotationInfo{
		Company: model.CompanyQuotationInfo{
			ID:          1,
			Company:     "บริษัท สยามชำนาญกิจ จำกัด",
			Branch:      "สำนักงานใหญ่",
			Address:     "เลขที่ 3 อาคารพร้อมพันธุ์ 3 ชั้น 10 ห้อง 1001 ซอยลาดพร้าว 3 ถนนลาดพร้าว แขวงจอมพล เขตจตุจักร กรุงเทพมหานคร",
			TaxID:       "0105556042151",
			PhoneNumber: "+66979515936",
		},
		Customer: model.CustomerQuotationInfo{
			ID:      3,
			Company: "บริษัท อยุธยา แคปปิตอล เซอร์วิสเซส จำกัด",
			Branch:  "สำนักงานใหญ่",
			Address: "อาคารกรุงศรีเพลินจิต ทาวเวอร์ 550 ถนนเพลินจิต แขวงเขตปทุมวัน กรุงเทพมหานคร",
			TaxID:   "0105537133562",
		},
		QuotationNumber: "QT256109-101002",
		Contact: model.Contact{
			Name:        "Nopparat Slisatkorn",
			Email:       "nopparat.slisatkorn@krungsri.com",
			PhoneNumber: "",
		},
		ProjectName: "โครงการ 3 Days Software Testing in Action Workshop",
		Orders: []model.Order{
			{
				OrderCourse:  "ค่าฝึกอบรม Software Testing in Action จำนวน 3 วัน วันพุธที่ 20 - วันศุกร์ที่ 22 มิถุนายน พ.ศ.2561",
				Amount:       1,
				PricePerUnit: "100,000.00",
				Price:        "100,000.00",
			},
		},
		Payment: model.Payment{
			TotalPrice:         "100,000.00",
			Discount:           "0.00",
			PriceAfterDiscount: "100,000.00",
			VAT:                "7,000.00",
			NetTotalPrice:      "107,000.00",
			TotalPriceThai:     "หนึ่งแสนเจ็ดพันบาทถ้วน",
		},
		IncludeVAT: false,
	}
	quotationForm := model.QuotationForm{
		CompanyTaxID:  "0105556042151",
		CustomerTaxID: "0105537133562",
		Contact: model.Contact{
			Name:        "Nopparat Slisatkorn",
			Email:       "nopparat.slisatkorn@krungsri.com",
			PhoneNumber: "",
		},
		ProjectName: "โครงการ 3 Days Software Testing in Action Workshop",
		Orders: []model.Order{
			{
				OrderCourse:  "ค่าฝึกอบรม Software Testing in Action จำนวน 3 วัน วันพุธที่ 20 - วันศุกร์ที่ 22 มิถุนายน พ.ศ.2561",
				Amount:       1,
				PricePerUnit: "100,000.00",
			},
		},
		IncludeVAT: false,
		Payment: model.Payment{
			Discount: "0.00",
		},
	}
	quotationService := QuotationServiceMySQL{
		QuotationRepository: &mockQuotationRepository{},
	}

	actualQuotationInfo, err := quotationService.CreateQuotation(quotationForm)

	if err != nil {
		t.Errorf("expect no error but it got %s", err)
	}
	if expectedQuotationInfo.Company != actualQuotationInfo.Company {
		t.Errorf("expect company is %v but it got %v", expectedQuotationInfo.Company, actualQuotationInfo.Company)
	}
	if expectedQuotationInfo.Customer != actualQuotationInfo.Customer {
		t.Errorf("expect customer is %v but it got %v", expectedQuotationInfo.Customer, actualQuotationInfo.Customer)
	}
	if expectedQuotationInfo.QuotationNumber != actualQuotationInfo.QuotationNumber {
		t.Errorf("expect quotation number is %s but it got %s", expectedQuotationInfo.QuotationNumber, actualQuotationInfo.QuotationNumber)
	}
	if expectedQuotationInfo.Contact != actualQuotationInfo.Contact {
		t.Errorf("expect contact is %v but it got %v", expectedQuotationInfo.Contact, actualQuotationInfo.Contact)
	}
	if expectedQuotationInfo.ProjectName != actualQuotationInfo.ProjectName {
		t.Errorf("expect project name is %s but it got %s", expectedQuotationInfo.ProjectName, actualQuotationInfo.ProjectName)
	}
	if expectedQuotationInfo.Orders[0] != actualQuotationInfo.Orders[0] {
		t.Errorf("expect first order is %v but it got %v", expectedQuotationInfo.Orders[0], actualQuotationInfo.Orders[0])
	}
	if expectedQuotationInfo.Payment != actualQuotationInfo.Payment {
		t.Errorf("expect payment is %v but it got %v", expectedQuotationInfo.Payment, actualQuotationInfo.Payment)
	}
}
