package main 
import "fmt"

func main() {
	/* ketika membuat constan nilai variablenya atau data harus di isi di awal */
	const firstName string = "diant"
	const lastName = "roys"
	const value = 500000000
	const age uint16 = 24 /* walaupun variable tidak digunakan dia tidak akan eror seperti var,  */
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	/* mendeklarasikan multiple variable constant  */
	const (
		asal string = "Petang"
		hoby 		= "Fishing"
		nilai int8 	= 100
	)
	fmt.Println(asal)
	fmt.Println(hoby)
	fmt.Println(nilai)
}