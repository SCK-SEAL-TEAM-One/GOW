package service_test

import (
	"gow/model"
	. "gow/service"
	"testing"
	"time"
)

func Test_CreateQuotation_Input_QuotationForm_Should_Be_Return_QuotationID_With_No_Error(t *testing.T) {
	expectedQuotationID := int64(1)
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
				PricePerUnit: 100000.00,
			},
		},
		IncludeVAT: false,
		Payment: model.Payment{
			Discount: 0.00,
		},
	}
	quotationService := QuotationServiceMySQL{
		QuotationRepository: &mockQuotationRepository{},
		OrderRepository:     &mockOrderRepository{},
		CustomerService:     &mockCustomerService{},
		CompanyService:      &mockCompanyService{},
		GetCurrentTime: func() time.Time {
			fixedTime, _ := time.Parse("2006-Jan-02", "2018-Apr")
			return fixedTime
		},
	}
	actualQuotationID, err := quotationService.CreateQuotation(quotationForm)

	if err != nil {
		t.Errorf("expect no error but it got %s", err)
	}
	if expectedQuotationID != actualQuotationID {
		t.Errorf("expect quotationID %d but it got %d", expectedQuotationID, actualQuotationID)
	}
}

func Test_GenerateQuotationNumber_Input_Year_2018_Month_4_QuotationNumber_2_Should_Be_QuotationNumber_QT256104_101002(t *testing.T) {
	expectedQuotationNumber := "QT256104-101002"
	quotationService := QuotationServiceMySQL{
		QuotationRepository: &mockQuotationRepository{},
		GetCurrentTime: func() time.Time {
			fixedTime, _ := time.Parse("2006-Jan", "2018-Apr")
			return fixedTime
		},
	}
	actualQuotationNumber := quotationService.GenerateQuotationNumber()

	if expectedQuotationNumber != actualQuotationNumber {
		t.Errorf("expect %s but it got %s", expectedQuotationNumber, actualQuotationNumber)
	}
}
