package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// address2 := address1 // copy value

	// address2.City = "Pulau Seribu"

	// fmt.Println(address1)
	// fmt.Println(address2)

	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := &address1 // pointer ke address1

	address2.City = "Kepulauan Seribu"
	fmt.Println(address1)
	fmt.Println(address2)
}