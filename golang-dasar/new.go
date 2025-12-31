package main

import "fmt"

type Address struct {
	City, Province, Country  string
}

func main() {
	var alamat1 *Address = new(Address)// membuat alamat baru di memory heap
	var alamat2 *Address = alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}