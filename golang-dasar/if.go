package main

import "fmt"

func main() {
	name := "Robert"

	if name == "Rafi" {
		fmt.Println("Hello Rafi")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hello, siapa kamu?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah sesuai")
	}
}