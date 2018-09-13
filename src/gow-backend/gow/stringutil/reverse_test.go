package stringutil_test

import (
	. "gow/stringutil"
	"testing"
)

func Test_Reverse_Input_100000_Should_Be_000001(t *testing.T) {
	expectedReversedString := "000001"
	numberString := "100000"

	actualReversedString := Reverse(numberString)

	if expectedReversedString != actualReversedString {
		t.Errorf("expect %s but it got %s", expectedReversedString, actualReversedString)
	}
}
