package student

import "fmt"

type Student struct {
	StdId      int
	StdName    string
	StdSurname string
	StdMidterm int
	StdFinal   int
}

func (s Student) CalculateExamAverage() {
	average := float32(s.StdMidterm)*0.4 + float32(s.StdFinal)*0.6

	fmt.Printf("%d %s %s average of the exams are = %.2f\n", s.StdId, s.StdName, s.StdSurname, average)

}
