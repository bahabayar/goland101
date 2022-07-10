package main

import "fmt"

func main() {
	var name = "Baha"
	var age int = 23
	var isMarried bool = false

	var weight float64 = 75.64
	var height int = 180

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)
	fmt.Println(height)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
	fmt.Printf("%T\n", weight)
	fmt.Printf("%T\n", height)
}
