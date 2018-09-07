package service_test

import (
	"gow/model"
	. "gow/service"
	"testing"
)

func Test_CreateQuotation_Input_QuotationForm_Should_Be_QuotationInfo(t *testing.T) {
	quotationService := QuotationServiceMySQL{
		QuotationRepository: &MockQuotationRepository{},
	}
	expectedQuotation := model.QuotationInfo{
		Company: model.CompanyQuotationInfo{
			ID:          3,
			Company:     "บริษัท ชู ฮา ริ จำกัด",
			Branch:      "สำนักงานใหญ่",
			Address:     "เลขที่ 3 อาคารพร้อมพันธ์ุ 3 ชั้น 10 ห้อง 1002 ซอยลาดพร้าม 3 ถนนลาดพร้าว แขวงจอมพล เขตจตุจักร กรุงเทพมหานคร 10900",
			TaxID:       "0105561001221",
			PhoneNumber: "+66979575936",
		},
		Customer: model.CustomerQuotationInfo{
			ID:      3,
			Company: "บริษัท อยุธยา แคปปิตอล เซอร์วิสเซส จำกัด",
			Branch:  "สำนักงานใหญ่",
			Address: "อาคารกรุงศรีเพลินจิต ทาวเวอร์ 550 ถนนเพลินจิต แขวงเขตปทุมวัน กรุงเทพมหานคร",
			TaxID:   "0105537133562",
		},
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
			Discount:           "0.00",
			PriceAfterDiscount: "100,000.00",
			VAT:                "7,000.00",
			TotalPrice:         "107,000.00",
			TotalPriceThai:     "หนึ่งแสนเจ็ดพันบาทถ้วน",
		},
		IncludeVAT: false,
	}
	quotationForm := model.QuotationForm{
		CompanyID:  2,
		CustomerID: 3,
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
	}

	createdQuotation, _ := quotationService.CreateQuotation(quotationForm)

	if expectedQuotation.Company != createdQuotation.Company {
		t.Errorf("expect quotation company %v but got it %v", expectedQuotation.Company, createdQuotation.Company)
	}
	if expectedQuotation.Customer != createdQuotation.Customer {
		t.Errorf("expect %v but got it %v", expectedQuotation.Customer, createdQuotation.Customer)
	}
	if expectedQuotation.Payment != createdQuotation.Payment {
		t.Errorf("expect %v but got it %v", expectedQuotation.Payment, createdQuotation.Payment)
	}
	if len(expectedQuotation.Orders) != len(createdQuotation.Orders) {
		t.Errorf("expect %v but got it %v", len(expectedQuotation.Orders), len(createdQuotation.Orders))
	}
}

func Test_CalculatePrice_Input_Amount_And_PricePerUnit_Should_Be_Price(t *testing.T) {
	expectedPrice := 100000.00
	amount := 1
	pricePerUnit := 100000.00

	actualPrice := CalculatePrice(amount, pricePerUnit)

	if expectedPrice != actualPrice {
		t.Errorf("expect %f but got %f", expectedPrice, actualPrice)
	}
}

func Test_CalculateDiscount_Input_Price_And_Discount_Should_Be_PriceAfterDiscount(t *testing.T) {
	expectedPriceAfterDiscount := 100000.00
	price := 100000.00
	discount := 0.00

	actualPriceAfterDiscount := CalculateDiscount(price, discount)

	if expectedPriceAfterDiscount != actualPriceAfterDiscount {
		t.Errorf("expect %f but got %f", expectedPriceAfterDiscount, actualPriceAfterDiscount)
	}
}

func Test_CalculateVat_Input_Price_And_VatRate_Should_Be_Vat(t *testing.T) {
	expectedVat := 7000.00
	price := 100000.00
	vatRate := 7.00

	actualVat := CalculateVat(price, vatRate)

	if expectedVat != actualVat {
		t.Errorf("expect %f but got %f", expectedVat, actualVat)
	}
}
