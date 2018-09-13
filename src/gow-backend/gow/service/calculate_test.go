package service_test

import (
	. "gow/service"
	"testing"
)

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

func Test_AddComma_Input_100000_00_Should_Be_100_comma_000_00(t *testing.T) {
	expectedPriceFormat := "100,000.00"
	price := 100000.00

	actualPriceFormat := AddComma(price)

	if expectedPriceFormat != actualPriceFormat {
		t.Errorf("expect %s but got %s", expectedPriceFormat, actualPriceFormat)
	}
}

func Test_CalculateNetTotalPrice_Input_PriceAfterDiscount_100000_And_VatFee_7000_Should_Be_NetTotalPrice_107000(t *testing.T) {
	expectedNetTotalPrice := 107000.00
	price := 100000.00
	vatFee := 7000.00

	actualNetTotalPrice := CalculateNetTotalPrice(price, vatFee)

	if expectedNetTotalPrice != actualNetTotalPrice {
		t.Errorf("expect %f but got %f", expectedNetTotalPrice, actualNetTotalPrice)
	}
}
