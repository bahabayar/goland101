package main

import (
	"fmt"
	"math"
)

func main() {

	//total := 22.50 / 2
	//fmt.Printf("%T, %v\n", total, total)
	//fmt.Printf("%T, %v\n", (9 / 3), (9 / 3))
	r := 3.0
	const piNumber = 3.14
	areaOfCircle := piNumber * (math.Pow(r, 2))
	fmt.Println(areaOfCircle)

}
