package main

import "fmt"

func main() {
	var nilaiAkhir = 100
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir >= 70
	var lulusAbsensi bool = absensi >= 75
	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println("Lulus:", lulus)
} 