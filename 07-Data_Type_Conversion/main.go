package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	x := 18
	y := 18.53

	someStringValue := "175"

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	fmt.Println(float64(x) + y)
	fmt.Println(x + int(y))

	intVal, err := strconv.Atoi(someStringValue)

	fmt.Println(intVal)
	fmt.Println(err)
	fmt.Println(reflect.TypeOf(intVal))
	fmt.Printf("%v %T\n", someStringValue, someStringValue)

}
