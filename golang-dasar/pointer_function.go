package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// var address *Address = new(Address)
	address := new(Address)
	ChangeCountryToIndonesia(address)

	fmt.Println(address)
}