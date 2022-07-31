package main

import "fmt"

type shape interface {
	area() int
	perimeter() int
}

type rectangle struct {
	longSide  int
	shortSide int
}
type square struct {
	side int
}

func (r rectangle) area() int {
	return r.shortSide * r.longSide
}
func (r rectangle) perimeter() int {
	return 2*r.longSide + 2*r.shortSide
}

func (s square) area() int {
	return s.side * s.side
}
func (s square) perimeter() int {
	return 4 * s.side
}

func main() {
	sqr := square{side: 10}
	rctngl := rectangle{longSide: 15, shortSide: 10}

	shapes := []shape{sqr, rctngl}

	calculateAll(shapes)

	fmt.Println(sqr.area())
	fmt.Println(sqr.perimeter())
	fmt.Println(rctngl.area())
	fmt.Println(rctngl.perimeter())

}

func calculateAll(s []shape) {
	for _, shp := range s {
		fmt.Println(shp.area())
		fmt.Println(shp.perimeter())
	}
}

func writeThemAll(value interface{}) {
	fmt.Printf("%T , %v", value, value)
}
