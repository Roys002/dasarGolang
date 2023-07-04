package main

import "fmt"

func main() {
	/* cara pertama, variable di tentukan type datannya */
	var name string
	name = "Royas"
	fmt.Println(name)
	name = "Mada"
	fmt.Println(name)
/* cara ke 2, type data tidak perlu di deklarasikan krna si golang tau type data tersebut */
	var data  = "Royas"
	fmt.Println(data)
	data = "Okeeh"
	fmt.Println(data)
	var age int8 = 25
	fmt.Println(age)
	/* tidak perlu menggunakan var saat membuat variable cukup dengan := */
	pekerjaan := "web Developer"
	fmt.Println(pekerjaan)
	dateArrival := " 15 agustus 2022"
	fmt.Println(dateArrival)
	/* membuat multiple variable */
	var (
		firstName = "Bobo"
		lastName = "hooo"
		umur int8 = 8
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(umur)
}