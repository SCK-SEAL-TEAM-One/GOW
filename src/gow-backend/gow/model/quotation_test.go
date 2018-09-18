package model_test

import (
	. "gow/model"
	"testing"
)

func Test_GetTotalPrice_With_Order_1_100000_Should_Be_100000(t *testing.T) {
	expectedTotalPrice := 100000.00
	quotationForm := QuotationForm{
		Orders: []Order{
			{
				OrderCourse:  "ค่าฝึกอบรม Software Testing in Action จำนวน 3 วัน วันพุธที่ 20 - วันศุกร์ที่ 22 มิถุนายน พ.ศ.2561",
				Amount:       1,
				PricePerUnit: 100000.00,
			},
		},
	}

	actualTotalPrice := quotationForm.GetTotalPrice()
	if expectedTotalPrice != actualTotalPrice {
		t.Errorf("expect %.3f but it got %.3f", expectedTotalPrice, actualTotalPrice)
	}
}

func Test_GetPrice_Order_1_100000_Should_Be_100000(t *testing.T) {
	expectedPrice := 100000.00
	order := Order{
		OrderCourse:  "ค่าฝึกอบรม Software Testing in Action จำนวน 3 วัน วันพุธที่ 20 - วันศุกร์ที่ 22 มิถุนายน พ.ศ.2561",
		Amount:       1,
		PricePerUnit: 100000.00,
	}

	actualPrice := order.GetPrice()
	if expectedPrice != actualPrice {
		t.Errorf("expect %.3f but it got %.3f", expectedPrice, actualPrice)
	}
}

func Test_GetPriceAfterDiscount_Order_1_100000_Discount_0_Should_Be_100000(t *testing.T) {
	expectedPriceAfterDiscount := 100000.00
	quotationForm := QuotationForm{
		Orders: []Order{
			{
				OrderCourse:  "ค่าฝึกอบรม Software Testing in Action จำนวน 3 วัน วันพุธที่ 20 - วันศุกร์ที่ 22 มิถุนายน พ.ศ.2561",
				Amount:       1,
				PricePerUnit: 100000.00,
			},
		},
	}

	actualPriceAfterDiscount := quotationForm.GetPriceAfterDiscount()
	if expectedPriceAfterDiscount != actualPriceAfterDiscount {
		t.Errorf("expect %.3f but it got %.3f", expectedPriceAfterDiscount, actualPriceAfterDiscount)
	}
}

func Test_GetVATFee_Order_1_100000_Discount_0_VAT_RATE_7_Should_Be_7000(t *testing.T) {
	expectedVatFee := 7000.00
	quotationForm := QuotationForm{
		Orders: []Order{
			{
				OrderCourse:  "ค่าฝึกอบรม Software Testing in Action จำนวน 3 วัน วันพุธที่ 20 - วันศุกร์ที่ 22 มิถุนายน พ.ศ.2561",
				Amount:       1,
				PricePerUnit: 100000.00,
			},
		},
		VATRate: 7.00,
	}

	actualVatFee := quotationForm.GetVATFee()
	if expectedVatFee != actualVatFee {
		t.Errorf("expect %.3f but it got %.3f", expectedVatFee, actualVatFee)
	}
}

func Test_GetNetTotalPrice_Order_1_100000_Discount_0_VAT_RATE_7_Should_Be_107000(t *testing.T) {
	expectedNetTotalPrice := 107000.00
	quotationForm := QuotationForm{
		Orders: []Order{
			{
				OrderCourse:  "ค่าฝึกอบรม Software Testing in Action จำนวน 3 วัน วันพุธที่ 20 - วันศุกร์ที่ 22 มิถุนายน พ.ศ.2561",
				Amount:       1,
				PricePerUnit: 100000.00,
			},
		},
		VATRate: 7.00,
	}

	actualNetTotalPrice := quotationForm.GetNetTotalPrice()
	if expectedNetTotalPrice != actualNetTotalPrice {
		t.Errorf("expect %.3f but it got %.3f", expectedNetTotalPrice, actualNetTotalPrice)
	}
}

func Test_GetTotalPriceThai_NetTotalPrice_107000_Should_Be_One_Hundred_Thousand_Seven_Thousand_Baht(t *testing.T) {
	expectedTotalPriceThai := "หนึ่งแสนเจ็ดพันบาทถ้วน"
	quotationForm := QuotationForm{
		Orders: []Order{
			{
				OrderCourse:  "ค่าฝึกอบรม Software Testing in Action จำนวน 3 วัน วันพุธที่ 20 - วันศุกร์ที่ 22 มิถุนายน พ.ศ.2561",
				Amount:       1,
				PricePerUnit: 100000.00,
			},
		},
		VATRate: 7.00,
	}

	actualTotalPriceThai := quotationForm.GetTotalPriceThai()
	if expectedTotalPriceThai != actualTotalPriceThai {
		t.Errorf("expect %s but it got %s", expectedTotalPriceThai, actualTotalPriceThai)
	}
}
