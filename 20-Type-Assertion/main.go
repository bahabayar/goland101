package main

import "fmt"

/*First Using */
func control(i interface{}) {
	value, control := i.(int)
	if control {
		fmt.Println("Type of value is integer", value)
	} else {
		fmt.Println("Type of value is not an integer")
	}
}

/*Second Using*/
func findType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type of parameter is int")
	case float64:
		fmt.Println("Type of parameter is float")
	case string:
		fmt.Println("Type of parameter is string")
	default:
		fmt.Println("Parameter has an unknown type ")
	}

}

func main() {

	number := 3
	float := 3.14
	word := "Test"
	control(number)
	control(float)
	control(word)

	findType(number)
	findType(float)
	findType(word)
	findType(nil)

}
