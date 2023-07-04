package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool
	/* type digunakan untuk membuat alias (maksud dari code di atas : 
		string digantikan dengan nama NoKtp sehingga, NoKtp itu sama dengan string) */
	var noKtpKu NoKtp = "25252525"
	var Kiku Married = true

	fmt.Println(noKtpKu)
	fmt.Println(Kiku)
}