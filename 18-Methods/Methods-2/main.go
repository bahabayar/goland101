package main

import "fmt"

type Student struct {
	stdNo      int
	stdName    string
	stdSurname string
	stdAge     int
}

func (s Student) showStudent() {
	fmt.Println(s.stdNo, s.stdName, s.stdSurname, s.stdAge)
}
func (s *Student) changeAge(newAge int) {
	s.stdAge = newAge
}

func main() {
	student := Student{stdNo: 123, stdName: "Baha", stdSurname: "Bayar", stdAge: 12}
	fmt.Println("Before the age is changed")
	fmt.Println("===========================")
	student.showStudent()
	fmt.Println("After the age is changed")
	fmt.Println("===========================")
	student.changeAge(23)
	student.showStudent()
}
