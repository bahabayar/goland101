package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//func firstRoutine(channel chan bool) {
//	fmt.Println("My first Routine")
//	channel <- true
//}

func sumOnlyOddNumbers(oddChannel chan int, number int) {
	total := 0
	if number != 0 {
		for i := 1; i <= number; i++ {
			if i%2 != 0 {
				total += i
			}
		}
	}
	oddChannel <- total

}
func sumOnlyEvenNumbers(evenChannel chan int, number int) {
	total := 0
	if number != 0 {
		for i := 1; i <= number; i++ {
			if i%2 == 0 {
				total += i
			}
		}
	}
	evenChannel <- total

}

func printY() {
	for i := 0; i < 100; i++ {
		fmt.Print("Y")
	}

}
func printX() {

	for i := 0; i < 100; i++ {
		fmt.Print("X")
	}
	wg.Done() /*Routine is done*/
}

func main() {

	wg.Add(1) /*This method used for to tell wait for other go routines for main routine*/
	go printX()
	wg.Wait()
	printY()
	//time.Sleep(time.Second)
	//var ch chan int
	//if ch == nil {
	//	ch = make(chan int)
	//	fmt.Printf("ch type is = %T\n", ch)
	////}
	//end := make(chan bool)
	//go firstRoutine(end)
	//<-end
	//fmt.Println("The End")

	//oChannel := make(chan int)
	//eChannel := make(chan int)
	//
	//go sumOnlyOddNumbers(oChannel, 10)
	//go sumOnlyEvenNumbers(eChannel, 10)
	//
	//totalOdds, totalEvens := <-oChannel, <-eChannel
	//fmt.Println("Odds Numbers + EvenNumbers = ", totalOdds+totalEvens)
}
