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
