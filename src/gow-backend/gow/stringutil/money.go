package stringutil

import (
	"fmt"
	"strings"
)

const (
	IntegerUnit = 0
	DecimalUnit = 1
)

func ConvertMoneyToThaiCharactor(number float64) string {
	numbrtString := fmt.Sprintf("%.2f", number)
	splitNumber := strings.Split(numbrtString, ".")
	numberInteger := splitNumber[IntegerUnit]
	numberDecimal := splitNumber[DecimalUnit]
	thaiCharactorInteger := ConvertNumberToThaiCharactor(numberInteger)
	if numberDecimal != "00" {
		thaiCharactorDecimal := ConvertNumberToThaiCharactor(numberDecimal)
		return fmt.Sprintf("%sบาท%sสตางค์", thaiCharactorInteger, thaiCharactorDecimal)
	}

	return fmt.Sprintf("%sบาทถ้วน", thaiCharactorInteger)

}

func ConvertNumberToThaiCharactor(numberString string) string {
	unitNumber := map[int]string{
		2: "สิบ",
		3: "ร้อย",
		4: "พัน",
		5: "หมื่น",
		6: "แสน",
		7: "ล้าน",
	}
	numberThai := map[string]string{
		"1": "หนึ่ง",
		"2": "สอง",
		"3": "สาม",
		"4": "สี่",
		"5": "ห้า",
		"6": "หก",
		"7": "เจ็ด",
		"8": "แปด",
		"9": "เก้า",
		"0": "",
	}
	var thaiCharactorInteger string
	lengthNumberInteger := len(numberString)
	for _, unitsInteger := range numberString {
		if string(unitsInteger) != "0" {
			thaiCharactorInteger += fmt.Sprintf("%s%s", numberThai[string(unitsInteger)], unitNumber[lengthNumberInteger])
		}
		lengthNumberInteger--
	}
	return thaiCharactorInteger
}
