package main

import (
	"fmt"
)

func main() {
	var name [3]string
	name[0] = "R"
	name[1] = "O"
	name[2] = "Y"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[1])

	var null  []string
	fmt.Println(len(null))
	fmt.Println(len(name))
	fmt.Println(len(values))
}