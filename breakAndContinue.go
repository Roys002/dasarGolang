package main

import (
"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("hasil perulangan break:", i)
	}

	for a := 0; a < 10; a++ {
		if a % 2 == 0 {
			continue
		} 
		fmt.Println("hasil perulangan continue:", a)
	}
}