package main

import "fmt"

func getFullName() (string, string) {
	return "Rafi", "Fadhlurrohman"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}