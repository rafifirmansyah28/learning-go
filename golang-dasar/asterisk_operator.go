package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address2 *Address = &address1 // pointer ke address1
	address2.City = "Kepulauan Seribu"

	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi Bandung

	*address2 = Address{"Bandung", "Jawa Barat", "Indonesia"} // mengubah pointer address2 dan address1 ke alamat baru
	fmt.Println(address1) // tetap
	fmt.Println(address2) // berubah menjadi Bandung
}