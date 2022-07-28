package main

import "github.com/meetthegolang/golang/18-Methods/Methods-4/nyp/student"

func main() {

	baha := student.Student{
		StdId:      515472,
		StdName:    "Baha",
		StdSurname: "Bayar",
		StdMidterm: 35,
		StdFinal:   75,
	}
	baha.CalculateExamAverage()
}
