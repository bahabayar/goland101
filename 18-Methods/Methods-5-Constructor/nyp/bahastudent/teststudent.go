package bahastudent

import "fmt"

type student struct {
	stdId      int
	stdName    string
	stdSurname string
	stdMidterm int
	stdFinal   int
}

func NewStudent(stdId int, stdName string, stdSurname string, stdMidterm int, stdFinal int) student {

	s := student{
		stdId:      3523,
		stdName:    "Baha",
		stdSurname: "Bayar",
		stdMidterm: 70,
		stdFinal:   80,
	}
	return s
}
func (s student) CalculateExamAverage() {
	average := float32(s.stdMidterm)*0.4 + float32(s.stdFinal)*0.6

	fmt.Printf("%d %s %s average of the exams are = %.2f\n", s.stdId, s.stdName, s.stdSurname, average)

}
