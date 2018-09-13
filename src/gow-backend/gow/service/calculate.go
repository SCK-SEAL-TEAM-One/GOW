package service

import (
	"bytes"
	"fmt"
	"gow/stringutil"
	"strings"
)

func CalculatePrice(amount int, pricePerUnit float64) float64 {
	price := float64(amount) * pricePerUnit
	return price
}

func CalculateDiscount(price, discount float64) float64 {
	priceAfterDiscount := price - discount
	return priceAfterDiscount
}

func CalculateVat(price, vatRate float64) float64 {
	vat := (price * vatRate) / InterestRate
	return vat
}

func AddComma(number float64) string {
	numberString := fmt.Sprintf("%.2f", number)
	splitedNumberString := strings.Split(numberString, ".")

	numberIntegerString := splitedNumberString[0]
	numberDecimalString := splitedNumberString[1]

	var buffer bytes.Buffer
	reverseNum := stringutil.Reverse(numberIntegerString)
	count := 0
	for index := 0; index < len(numberIntegerString); index++ {
		count++
		buffer.WriteString(reverseNum[index : index+1])
		if count == 3 && index != len(numberIntegerString)-1 {
			buffer.WriteString(",")
			count = 0
		}
	}

	numberIntegerString = stringutil.Reverse(buffer.String())
	return fmt.Sprintf("%s.%s", numberIntegerString, numberDecimalString)
}
