package stringutil_test

import (
	. "gow/stringutil"
	"testing"
)

func Test_AddComma_Input_100000_00_Should_Be_100_comma_000_00(t *testing.T) {
	expectedPriceFormat := "100,000.00"
	price := 100000.00

	actualPriceFormat := AddComma(price)

	if expectedPriceFormat != actualPriceFormat {
		t.Errorf("expect %s but got %s", expectedPriceFormat, actualPriceFormat)
	}
}
