package main

import (
	"Amit16110/Phone-number-cleanup/phone"
	"fmt"
)

func main() {

	num := phone.Number("+44 123 456789")
	format := phone.Format("+1 (613)-995-0253")
	areaCode := phone.AreaCode("+1 (613)-995-0253")
	// Print
	fmt.Println("Number:", num)
	fmt.Println("Format:", format)
	fmt.Println("Area code:", areaCode)

}
