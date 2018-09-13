package service

import (
	"bytes"
	"fmt"
	"gow/model"
	"gow/stringutil"
	"strconv"
	"strings"
)

const (
	InterestRate = 100
	IntegerUnit  = 0
	DecimalUnit  = 1
)

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

func CalculateNetTotalPrice(totalPrice, vatFee float64) float64 {
	return totalPrice + vatFee
}

func CalculateOrdersPrice(orders *[]model.Order) (float64, error) {
	var totalPrice float64
	for index, order := range *orders {
		priceWithoutComma := strings.Replace(order.PricePerUnit, ",", "", -1)
		pricePerUnit, err := strconv.ParseFloat(priceWithoutComma, 64)
		if err != nil {
			return 0.00, err
		}
		price := CalculatePrice(order.Amount, pricePerUnit)
		(*orders)[index].Price = AddComma(price)
		totalPrice += price
	}
	return totalPrice, nil
}
