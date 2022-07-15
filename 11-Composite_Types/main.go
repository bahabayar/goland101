package main

import "fmt"

func main() {

	var array = [...]bool{3: false, 5: true}
	fmt.Println(len(array))
	fmt.Println(array)

	names := [...]string{"Baha", "Reha", "Can"}

	fmt.Println(names[0])

	numbers := [5]int{25, 50, 75, 30, 45}
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}

	println(total)

	languages := [3]string{"Go", "C#", "Java"}
	numbers := [4]float64{12.5, 10.0, 58.9, 56.7}
	t_f := [5]bool{true, false, true, true, false}

	for index := range languages {
		fmt.Print("Index =", index, " ")
	}
	fmt.Println()
	for index, value := range numbers {
		fmt.Print("Index = ", index, "Value = ", value, "-")
	}
	fmt.Println()
	for _, value := range t_f {
		fmt.Print("Value = ", value, " ")
	}

	studentsNumber := [7]int{100, 101, 102, 103, 104, 105, 106}
	studentsGrades := [7]int{86, 26, 52, 100, 15, 5, 25}
	average := 0
	total := 0

	for _, value := range studentsGrades {
		total += value
	}
	average = total / len(studentsGrades)
	fmt.Println("Exam average is = ", average)

	for i := 0; i < len(studentsGrades); i++ {
		if studentsGrades[i] > average {
			fmt.Println("Student Number : ", studentsNumber[i], "is above average. Exam score is ", studentsGrades[i])
		}
	}

	brands := [3][2]string{{"Google", "Yandex"}, {"Apple", "Samsung"}, {"Facebook", "Instagram"}}

	for i := 0; i < len(brands); i++ {
		fmt.Println()
		for j := 0; j < len(brands[i]); j++ {
			fmt.Println("i = ", i, "j =", j, "brands =", brands[i][j])
		}
	}

	//brands := [3][2]string{{"Google", "Yandex"}, {"Apple", "Samsung"}, {"Facebook", "Instagram"}}
	//
	//for _, i := range brands {
	//	fmt.Println()
	//	for _, j := range i {
	//		fmt.Println(j)
	//	}
	//}

	//	array1 := [...]float64{10.2, 54.6, 96.5}
	//	x := Sum(array1)
	//	fmt.Printf("Your array sum is = %v", x)
	//}
	//
	//

	//Sum /*Using Pointer*/
	array1 := [...]float64{10.2, 54.6, 96.5}
	x := Sum(&array1)
	fmt.Printf("Your array sum is = %v", x)
}

//func Sum(array [3]float64) (total float64) {
//	for _, v := range array {
//		total += v
//	}
//	return

// Sum /*Using Pointer*/
func Sum(array *[3]float64) (total float64) {
	for _, v := range array {
		total += v
	}
	return
}
