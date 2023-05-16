package phone

import (
	"fmt"
	"regexp"
	"strings"
)

func Number(phoneNumber string) string {
	re := regexp.MustCompile(`\D`)
	cleanNumber := re.ReplaceAllString(phoneNumber, "")

	if len(cleanNumber) > 10 && strings.HasPrefix(cleanNumber, "1") {
		cleanNumber = cleanNumber[1:]
	}

	if len(cleanNumber) != 10 {
		return ""
	}

	return cleanNumber
}

func Format(phoneNumer string) string {
	cleanNumber := Number(phoneNumer)

	if cleanNumber == "" {
		return ""
	}

	formattedNumber := fmt.Sprintf("(%s) %s-%s", cleanNumber[0:3], cleanNumber[3:6], cleanNumber[6:10])

	return formattedNumber
}

func AreaCode(phoneNumber string) string {
	cleanNumber := Number(phoneNumber)

	if cleanNumber == "" {
		return ""
	}

	areaCode := cleanNumber[0:3]

	return areaCode
}
