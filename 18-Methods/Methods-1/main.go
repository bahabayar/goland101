package main

import "fmt"

type Student struct {
	stdId      int
	stdName    string
	stdSurname string
}
type Rectangle struct {
	longSide, shortSide int
}
type Square struct {
	side int
}

func main() {
	student := Student{
		stdId:      456489,
		stdName:    "Baha",
		stdSurname: "Bayar",
	}
	student.showStudent()
	rectangle := Rectangle{longSide: 15, shortSide: 10}
	square := Square{side: 4}
	fmt.Println("Area of Rectangle = ", rectangle.findArea())
	fmt.Println("Area of Square = ", square.findArea())
}

func (s Student) showStudent() {
	fmt.Println(s.stdId, s.stdName, s.stdSurname)
}

func (r Rectangle) findArea() int {
	return r.longSide * r.shortSide
}
func (s Square) findArea() int {
	return s.side * s.side
}
