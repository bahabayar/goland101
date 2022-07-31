package statistics

import (
	"fmt"
	"math"
)

func GetAvg(values []float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total / float64(len(values))
}

// GetStdev Calculate standard deviation
func GetStdev(values []float64) float64 {
	result := 0.0
	avg := GetAvg(values)
	total := 0.0
	for _, value := range values {

		total += math.Pow((value - avg), 2)
	}
	result = math.Sqrt(total / float64(len(values)))
	return result
}

func GetVariance(values []float64) float64 {
	return math.Pow(GetStdev(values), 2)
}

func init() {
	fmt.Println("Package Loaded!")
}
