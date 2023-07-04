package main
import (
	"fmt"
)

func main() {
	var month = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
	var slice1 = month[4:8]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = month[9:]
	fmt.Println(slice2)

	/* append digunakan untuk mengubah value dalam array tampa mengubah parent arraysnya 
		dan juga jika array sudah penuh dia akan membuar array baru dengan nama diant, hanya 
		berlaku jika capasity dari arraynya sudah penuh, jika masih ada ruang/capasity masih ada
		berubahan data akan berimpack ke parentnya 
	*/
	var slice3 = append(slice2, "Diant")
	fmt.Println(slice3)
	slice3[2] = "NotDecember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(month)

	/* lebih aman menggunakan make karena arraynya langsung dimasukkan ke newSlice, angka 2 itu panhang arraynya
	sedangkan 5 itu capasity nya */
	newSlice := make([]string , 2,5)
	newSlice[0] = "Boys"
	newSlice[1] = "Girl"

	fmt.Println(newSlice,"newSlice")
	fmt.Println(len(newSlice), "length NewSlice")
	fmt.Println(cap(newSlice), "Capacity NewSlice")

	// code di bawah untuk copy slice, panjang dan capasity harus sama dengan sumber copy nya
	copySlice := make([]string, len(newSlice), cap(newSlice))
	/* code di bawah "copySlice" merupkan destinationnya sedangkan newSlice adalah sumber copynya */
	copy(copySlice, newSlice)
	fmt.Println(copySlice,"Copy slice dari newSlice")

	iniArray := [...]int {1,2,3,4,5}
	iniSlice := []int {1,2,3,4,5,}
	fmt.Println(iniArray,"ini Array")
	fmt.Println(iniSlice, "ini slice")
}