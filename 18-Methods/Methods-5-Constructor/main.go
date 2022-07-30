package main

import "github.com/meetthegolang/golang/18-Methods/Methods-5-Constructor/nyp/bahastudent"

func main() {

	std := bahastudent.NewStudent(10, "Baha", "Bayar", 56, 100)
	std.CalculateExamAverage()
}
