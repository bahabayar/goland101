// Package clause
package main

// Import statement
import "fmt"

// My Code
func main() {
	/*
		var name string //Variable Declaration
		name = "Baha"   //Variable assign

		var name string = "Baha"
		var age int = 22
		var isMarried bool = false

		age := 22
		age = 23


		var firstName, lastName string = "Baha", "Bayar"  Multiple declaration and assign
	*/

	name := "Baha"
	age := 22
	isMarried := false
	age = 23

	fmt.Println("Hello", name)
	fmt.Println("Your age is", age)
	if isMarried {
		fmt.Println("You are married")
	} else {
		fmt.Println("You are not married")
	}

}
