package main

import "fmt"

type student struct {
	stdId      int
	stdName    string
	stdSurname string
}

func (s student) showStudent() string {
	return fmt.Sprintf("%d , %s %s", s.stdId, s.stdName, s.stdSurname)
}

type studentLessons struct {
	lessonName string
	midterm    int
	final      int
	student
}

func (sl studentLessons) showStudentExamResult() {
	fmt.Println("Student = ", sl.showStudent())
	fmt.Println("Lesson = ", sl.lessonName)
	fmt.Println(" Midterm = ", sl.midterm)
	fmt.Println("Final = ", sl.final)
	fmt.Println("Average = ", float32(sl.midterm)*0.4+float32(sl.final)*0.6)
}
func main() {
	student := student{
		stdId:      123,
		stdName:    "Baha",
		stdSurname: "Bayar",
	}
	lesson := studentLessons{
		lessonName: "Math",
		midterm:    70,
		final:      40,
		student:    student,
	}
	lesson.showStudentExamResult()
}
