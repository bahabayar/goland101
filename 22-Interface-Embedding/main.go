package main

import (
	"fmt"
	"math"
)

type calculations interface {
	average
	minimum
}
type average interface {
	findAvg() float64
}
type minimum interface {
	findMin() float64
}

type student struct {
	stdId      int
	stdName    string
	stdSurname string
	stdMidterm int
	stdFinal   int
}

func (s student) findAvg() float64 {
	return float64(s.stdMidterm)*0.40 + float64(s.stdFinal)*0.60
}
func (s student) findMin() float64 {
	minExam := (60 - float64(s.stdMidterm)*0.4) / 0.6
	return math.Round(minExam)
}

func main() {
	std := student{
		stdId:      123123,
		stdName:    "Baha",
		stdSurname: "Bayar",
		stdMidterm: 35,
		stdFinal:   100,
	}

	var clct calculations = std

	fmt.Println(clct.findMin())
	fmt.Println(clct.findAvg())
}
