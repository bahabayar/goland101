package main

import "fmt"

func main() {

	var slice1 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := []bool{false, true, true, false, true, false}

	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	fmt.Println(&slice1)
	fmt.Println(len(slice1))
	fmt.Println(slice2)

	///*Another way to using slice*/
	s := make([]int, 0, 3)
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("cap %d , len %d\n", cap(s), len(s))
	}

	/*Array To Slice*/

	var array1 = [5]string{"Go", "Java", "C#", "Python", "C++"}

	var slice3_ []string = array1[2:4]
	var slice4 []string = array1[1:3]

	fmt.Println("Array : ", array1)
	fmt.Println("Slice1 : ", slice3_)
	fmt.Println("Slice2", slice4)

	/*Copy Function*/
	mainSlice := []string{"Go", "Java", "C#", "Python", "C++"}
	targetSlice := make([]string, 2)
	numberOfCopies := copy(targetSlice, mainSlice)

	fmt.Println("Main Slice :", mainSlice)
	fmt.Println("Target Slice", targetSlice)
	fmt.Println("Number Of Copies", numberOfCopies)

	/*Append Function*/
	d := make([]int, 0, 3)
	a := make([]int, 2, 5)

	for i := 0; i < 5; i++ {
		d = append(d, i)
		fmt.Printf("First slice = cap %d , len %d , point %p \n", cap(d), len(d), d)
	}
	for i := 0; i < 5; i++ {
		a = append(a, i)
		fmt.Printf("Second slice = cap %d , len %d , point %p \n", cap(a), len(a), a)
	}

	/*Slider With Capitals*/
	Capitals()

	/*Slider With Using Parameter Function*/
	var numberArray = [4]float64{2.25, 58.48, 45.75, 78.50}
	fmt.Println(Sum(numberArray[:]))

}

func Capitals() {
	capitals := []string{"Ankara", "Madrid", "Paris", "Roma"}
	countries := []string{"Turkey", "Spain", "France", "Italy"}
	phoneCode := []string{"+90", "+34", "+33", "+39"}

	for i := 0; i < cap(capitals); i++ {
		fmt.Println("The capital of", countries[i], "is", capitals[i])
	}
	fmt.Println()
	for ind, val := range phoneCode {
		fmt.Println("The phone code of", countries[ind], "is", val)
	}
}

func Sum(a []float64) float64 {

	s := 0.0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
