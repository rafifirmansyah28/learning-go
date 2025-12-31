package main

import "fmt"

func main() {
	name := "Rafi"
	
	switch name {
		case "Rafi":
			fmt.Println("Hello Rafi")
		case "Budi":
			fmt.Println("Hello Budi")
		default:
			fmt.Println("Hello, siapa kamu?")
	}

	switch length := len(name); length > 5 {
		case true:
			fmt.Println("Nama terlalu panjang")
		case false:
			fmt.Println("Nama sudah sesuai")
	}

	otherName := "Budi"
	switch {
		case otherName == "Rafi":
			fmt.Println("Hello Rafi")
		case otherName == "Budi":
			fmt.Println("Hello Budi")
		default:
			fmt.Println("Hello, siapa kamu?")
	}
}