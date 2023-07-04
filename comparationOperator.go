package main

import "fmt"

func main() {
	var (
		value1 = 5
		value2 = 6
		name1 = "Roys"
		name2 = "Roys"

		result bool = value1 > value2
		result2 bool = value1 != value2
		result3 bool = name1 == name2
	)

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
}