package main

import (
	"fmt"
	"strconv"
)

func main() {
	//
	//v := "10"
	//fmt.Printf("%T ,%v \n", v, v)
	//number, _ := strconv.Atoi(v)
	//fmt.Printf("%T ,%v \n", number, number)

	var numb int = 34
	fmt.Printf("%T ,%v \n", numb, numb)

	n := strconv.Itoa(numb)

	fmt.Printf("%T ,%v \n", n, n)

	b, _ := strconv.ParseBool("true")
	fmt.Printf("%T ,%v \n", b, b)

}
