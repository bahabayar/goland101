package main

import (
	"fmt"
	"time"
)

func main() {

	switch a := 10; {
	case a == 10:
		fmt.Println("First Case is works!")
		fallthrough
	case a%2 == 0:
		fmt.Println("Second Case also works!")
	case a%5 != 0:
		fmt.Println("Third Case also works!")

		fmt.Print("Enter number : ")
		var n int
		fmt.Scanln(&n)
		if n%2 == 0 {
			fmt.Println(n, "is Even number")
		} else {
			fmt.Println(n, "is Odd number")
		}

		today := time.Now()
		t := today.Day()
		switch t {
		case 1, 28, 29, 30, 31:
			fmt.Println("It's time to pay dues")
		case 12, 13, 14, 15, 16, 17, 18:
			fmt.Println("It's time to learn Golang")
		default:
			fmt.Println("No information available for that day")
		}

		for i := 0; i < 10; i++ {
			fmt.Printf("%T ,%v\n", i, i)

			fmt.Print("Enter a number : ")
			var n int
			fmt.Scanln(&n)
			for a := 0; a < 10; a++ {
				total := n * a
				fmt.Printf("%v * %v = %v", n, a, total)
			}

		}
	}

	word := "Welcome To The Go Programming"

	for ind, c := range word {
		fmt.Printf("index is %d character is  = %c\n", ind, c)
	}

	fmt.Print("Enter number : ")
	var n int
	fmt.Scanln(&n)
	isPrimeNumber := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			isPrimeNumber = false
			break
		}
	}

	if isPrimeNumber {
		fmt.Printf("%v is Prime Number ", n)
	} else {
		fmt.Printf("%v is Not a Prime Number ", n)
	}

	for i := 0; i < 5; i++ {
		if i == 2 || i == 4 {
			continue
		}
		fmt.Println(i)
	}
}
