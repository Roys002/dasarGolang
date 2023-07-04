package main

import "fmt"

func main() {
	var (
		nilaiAkhir   = 80
		absensiAkhir = 90

		lulusNilaiAkhir   = nilaiAkhir >= 80
		lulusAbsensiAkhir = absensiAkhir >= 80
		lulus             = lulusAbsensiAkhir && lulusNilaiAkhir
	)
	fmt.Println(lulus)
}