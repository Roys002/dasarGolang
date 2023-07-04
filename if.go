package main 

import (
	"fmt"
)

func main() {
	var name = "Bob"
	if name == "Diant" {
		fmt.Println("Hallo Dian")
	}else if name == "Bob" {
		fmt.Println("Hallo Bob")
	}else {
		fmt.Println("Kau sape")
	}

	/* sort statment pada if, if di golang dapat mendeklarasikan sebuah variabel
	langsung di dalam if nya ,contoh pada variable length, dan ";" digunakan untuk
	memisahkan declarasi dengan kondisi nya  */
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("name anda pas")
	}
}