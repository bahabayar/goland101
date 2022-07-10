package main

import "fmt"

func main() {
	var name = "Baha"
	var age int = 23
	var isMarried bool = false

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
}
