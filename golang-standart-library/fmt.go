package main

import "fmt"

func main() {
	firstName := "Rafi"
	lastName := "Maulana"

	fmt.Println("Hello '", firstName, lastName, "'")
	
	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}