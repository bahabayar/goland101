package main

import "fmt"

var packVar = "Package Scope"

func main() {
	const birthday string = "16/10/1999"

	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	}

	const (
		a = iota + 100
		b
		c
	)
	const (
		d = iota + 2
		_
		e
		_
		f
	)

	funcVar := "Func Scope"

	fmt.Println(funcVar)
	fmt.Println(packVar)
	fmt.Println(birthday)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
}
