package main

import "fmt"

func main() {
	/* konvesi type data */
	var nilai32 int32 = 3320
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	/* 
	ketika ingin mengkonversi type data dari besar ke keci
	(int64 ke int16) akan terjadi perubahan nilai karena 
	nilai yang besar tidak bisa di handle, int16 akan balik ke 
	nilai minimumnya
	*/
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Diant" 
	var D = name[0] /* uint itu sama dengan byte */
	var dString string = string(D)
	
	fmt.Println(name, "name")
	fmt.Println(D, "D")
	fmt.Println(dString, "dString")
}