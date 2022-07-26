package main

import (
	"fmt"
	"strings"
)

func main() {

	stringValue := "Lets learn golang"
	fmt.Println(stringValue)
	makeStringUpperCase(&stringValue)
	fmt.Println(stringValue)

}

func practice1() {
	b := 255

	var a *int = &b

	b += 10
	fmt.Println("Variable value is = ", b)
	*a *= 2
	fmt.Println("Variable value is = ", *a)

}

func makeStringUpperCase(val *string) {
	(*val) = strings.ToUpper(*val)
}
