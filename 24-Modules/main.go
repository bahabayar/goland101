package main

import (
	"fmt"
	stat "github.com/meetthegolang/golang/24-Modules/statistics"
	"math/rand"
)

func main() {
	array := []float64{2, 1, 5, 4, 5, 1, 2, 3, 5}
	fmt.Println("Avg =", stat.GetAvg(array))
	fmt.Println("Stdev = ", stat.GetStdev(array))
	fmt.Println("Varyans = ", stat.GetVariance(array))

	fmt.Println(rand.Float64())
	fmt.Println(rand.Intn(100))
}
