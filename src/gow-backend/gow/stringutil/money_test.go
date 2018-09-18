package stringutil_test

import (
	. "gow/stringutil"
	"testing"
)

func Test_GetTotalPriceThai_Money_107000_Should_Be_One_Hundred_Thousand_Seven_Thousand_Baht(t *testing.T) {
	expectedTotalPriceThai := "หนึ่งแสนเจ็ดพันบาทถ้วน"
	money := 107000.00
	actualTotalPriceThai := ConvertMoneyToThaiCharactor(money)
	if expectedTotalPriceThai != actualTotalPriceThai {
		t.Errorf("expect %s but it got %s", expectedTotalPriceThai, actualTotalPriceThai)
	}
}

func Test_ConvertNumberToThaiCharactor_NumberString_107000_Should_Be_One_Hundred_Thousand_Seven_Thousand(t *testing.T) {
	expectedTotalPriceThai := "หนึ่งแสนเจ็ดพัน"
	money := "107000"
	actualTotalPriceThai := ConvertNumberToThaiCharactor(money)
	if expectedTotalPriceThai != actualTotalPriceThai {
		t.Errorf("expect %s but it got %s", expectedTotalPriceThai, actualTotalPriceThai)
	}
}
