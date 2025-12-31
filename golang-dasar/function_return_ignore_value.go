package main

import "fmt"

func getFullName() (string, string) {
	return "Rafi", "Fadhlurrohman"
}

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
}