package main

import "fmt"

func main() {

	//var factorialNumber int

	//compareNumbers(5, 2)
	//compareNumbers(3, 3)
	//println("Please enter a number for find the Factorial of a Number :")
	//fmt.Scan(&factorialNumber)
	//
	//fmt.Println(foundFactorialOfNumber(factorialNumber))

	//fmt.Println(sumNumbersRecursiveFunction(5))
	//letsLearnDefer(1, 2, 3, 4, 5, 6)

	a, b := 2, 5
	fmt.Println("Values before the calling callByRef Function")
	fmt.Printf("a=%v , b =%v\n", a, b)
	callByRef(&a, &b)
	fmt.Println("Values after the calling callByRef Function")
	fmt.Printf("a=%v , b =%v", a, b)

	/*CallingRefFunc*/
}

func letsLearnDefer(numbers ...int) {
	defer fmt.Println("Writing job is done!")
	for _, number := range numbers {
		defer fmt.Println(number)
	}
	fmt.Println("Writing job is started")
}
func callByRef(a *int, b *int) {
	c := *a
	*a = *b
	*b = c

}

func sumNumbersRecursiveFunction(number int) int {
	if number == 0 {
		return 0
	}
	return number + sumNumbersRecursiveFunction(number-1)
}

func compareNumbers(a int, b int) string {
	result := ""
	if a > b {
		result = "a is greater than b"
	} else if b > a {
		result = " b is greater than a"
	} else {
		result = "numbers are equal"
	}
	return result
}

func foundFactorialOfNumber(number int) int {
	result := 1
	for i := 1; i <= number; i++ {
		result *= i
	}
	return result
}

func anonymousFunctions() {
	functions := func() {
		fmt.Println("Hello , ı am a Anonymous Functions")
	}
	functions()
	fmt.Printf("%T\n", functions)

}

func closuresFunctions() {
	i := 0
	func() {
		fmt.Printf("i = %d\n", i)
	}()
}

func highOrderFunctions() {
	function := func(a, b int) int {
		return a % b
	}
	calculateModulo(function)
}

func calculateModulo(fnc func(a, b int) int) {
	fmt.Println("Result = ", fnc(5, 2))

}
