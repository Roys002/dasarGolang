package main

import (
"fmt"
)

func main() {
	name := "Diantee"

	switch name {
	case "Dian":
		fmt.Println("Hello Dian")
	case "Roys" :
		fmt.Println("Hello Roys")
	case "May" : 
		fmt.Println("Hello May")
	default:
		fmt.Println("Kau ni ape")
	}

	// short statment using switch
	switch length := len(name); length > 5 {
		case true:
			fmt.Println("Nama terlalu panjang")
		case false: 
			fmt.Println("Nama Cukup")
	}

	// di golang dapat membuat switch tanpa kondisi di depan swwitch tapi langsung di case
	lengt := len(name)
	switch {
	case lengt > 10 :
		fmt.Println("Nama Terlalu Panjang")
	case lengt > 5 :
		fmt.Println("Nama Cukup Panjang")
	default :
		fmt.Println("Nama Pas")
	}
}