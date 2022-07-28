package main

import "fmt"

type myFloat float32

func (a myFloat) division(b myFloat) myFloat {
	return a / b

}

func main() {
	a := myFloat(42)
	b := myFloat(7)
	fmt.Println(a.division(b))
	fmt.Println(b.division(a))
}
