package stringutil

import (
	"bytes"
	"fmt"
	"strings"
)

func AddComma(number float64) string {
	numberString := fmt.Sprintf("%.2f", number)
	splitedNumberString := strings.Split(numberString, ".")

	numberIntegerString := splitedNumberString[0]
	numberDecimalString := splitedNumberString[1]

	var buffer bytes.Buffer
	reverseNum := Reverse(numberIntegerString)
	count := 0
	for index := 0; index < len(numberIntegerString); index++ {
		count++
		buffer.WriteString(reverseNum[index : index+1])
		if count == 3 && index != len(numberIntegerString)-1 {
			buffer.WriteString(",")
			count = 0
		}
	}

	numberIntegerString = Reverse(buffer.String())
	return fmt.Sprintf("%s.%s", numberIntegerString, numberDecimalString)
}
