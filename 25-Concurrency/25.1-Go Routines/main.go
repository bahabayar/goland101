package main

import (
	"fmt"
	"time"
)

func myFirstRoutine() {
	fmt.Println("My First Routine")
}

func printNumbers() {
	for i := 1; i < 6; i++ {
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go myFirstRoutine()
	time.Sleep(1 * time.Second)
	fmt.Println("End")
}
