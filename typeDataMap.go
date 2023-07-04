package main

import (
	"fmt"
)

func main(){
	/* map[string] => string di dalamnya itu adalah key nya, makanya name di 
	bawah harus type data string, sedangkan string di luar map[string]
	merupakan value nya */
	// declarasi pendek
	person := map[string]string{
		"name": "Bob",
		"age": "20",
	}
	// menambahkan data
	person["gender"] = "male"
	fmt.Println(person)
	fmt.Println(person["name"],"name")
	fmt.Println(person["age"], "age")
	fmt.Println(person["gender"], "gender")

	// declarasi panjang
	var data  map[string]string = map[string]string {
		"Pertama": "Hallo",
		"Kedua": "Word",
	}
	fmt.Println(data)
	fmt.Println(data["Pertama"],"Pertama")
	fmt.Println(data["Kedua"], "Kedua")

	// function map
	book := make(map[string]string)
	book["title"] = "belajar dasar Golang"
	book["author"] = "Bob"
	book["job"] = "Programmer"
	book["wrong"] = "Ohhh"

	fmt.Println(book,"book3")
	// penggunaan delete (nama map, key)
	delete(book, "wrong")
	fmt.Println(book,"book2")
}