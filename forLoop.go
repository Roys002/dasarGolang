package main

import (
"fmt"
)

func main() {
	/* di golang untuk perulangan hanya ada for */
	counter := 1
	for counter <= 10 {
		fmt.Println("COunter Ke",counter)
		counter++
	}

	/* dataInt := 1 artinya init statemnt, di declarasi kan diawal
		dataInt <= 10 merupakan kondisi nya
		dataInt++ merupakan Post statment, artinya dataInt++ akan di jalankan di akhir perulangan */
	for dataInt := 1; dataInt <= 10; dataInt++{
		fmt.Printf("%d\n",dataInt);
	}

	/* for range, for bisa digunakan untuk iterasi terhadap semua data colection
		seperti array, slice, map
	*/
	// contoh pengambilan data secara manual dengan for menggunakan index nya (i)
	slice := []string{"Diant","Roy","Mada"}
	for i := 0 ; i < len(slice); i++{
		fmt.Println(slice[i])
	}

	// penggunaan range
	name := []string{"Diant","Roy","Mada"}
// index merupakan index nya dan value merupakan nilainya
	for index,value := range name {
		fmt.Println("Index",index ,"=", value)
	}
	// gunakan _ jika variable index tidak di perlukan/ kita hanya ingin mengambil value nya saja
	for _,value := range name {
		fmt.Println("Value nya adalah", value)
	}
/* penggunaan map pada for, */
	person := make(map[string]string)
	person["name"] = "Diant"
	person["age"] =  "25 years old"
	person["gender"] = "male"

	for key, value := range person {
		fmt.Println(key ,"=" , value)
	}
}