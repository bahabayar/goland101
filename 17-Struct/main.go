package main

import "fmt"

type Person struct {
	name, surname string
	age           int
}
type AnonymousPerson struct {
	string
	int
}
type Rectangle struct {
	longSide, shortSide int
}
type Student struct {
	stdNo      int
	stdName    string
	stdSurname string
}

func main() {
	createPerson()
	createRectangle()
	anonymousStruct()
	anonymousStructFields()
	pointerStructPractice()
	usingNewKeyWord()
}

func createPerson() {
	Baha := Person{
		name:    "Baha",
		surname: "Bayar",
		age:     23,
	}

	fmt.Println(Baha)
}
func createRectangle() {
	firstRectangle := Rectangle{longSide: 15, shortSide: 10}
	secondRectangle := Rectangle{longSide: 13, shortSide: 10}
	fmt.Println("First Rectangle = ", firstRectangle)
	fmt.Println("Second Rectangle = ", secondRectangle)
	secondRectangle.longSide = 20
	fmt.Println("First Rectangle = ", firstRectangle)
	fmt.Println("Second Rectangle = ", secondRectangle)
}
func anonymousStruct() {
	person := struct {
		name, surname string
		age           int
	}{name: "Mahmut Baha", surname: "Bayar", age: 23}
	fmt.Println(person)
}

func anonymousStructFields() {
	person := AnonymousPerson{"Baha", 23}
	fmt.Println("This anonymousStructFields =", person)
}

func pointerStructPractice() {
	student := Student{
		stdNo:      11235,
		stdName:    "Baha",
		stdSurname: "Bayar",
	}
	pointer := &student
	fmt.Println((*pointer).stdName)
	fmt.Println(student.stdName)
}

func usingNewKeyWord() {
	person := new(Person)
	person.name = "Baha"
	person.surname = "Bayar"
	person.age = 23
	fmt.Println("Name =", person.name)
	fmt.Println("Surname =", person.surname)
	fmt.Println("Age =", person.age)
	fmt.Println("PersonStruct =", person)
}
