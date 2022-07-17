package main

import "fmt"

func main() {

	var map1 map[string]int
	if map1 == nil {
		fmt.Println("Map1 is a empty ")
	}

	map2 := map[string]int{"Go": 1, "C#": 2, "Python": 3}

	fmt.Println(map2["Go"])
	fmt.Println(map2["C#"])

	/* Sending Map As Parameter To Function*/

	var map3 = make(map[string]string)
	map3["Go"] = "Lang"
	PrintValues(map3)
}

func PrintValues(map1 map[string]string) {
	fmt.Println(map1["Go"])
}
