package main

import "fmt"

func main() {
	var language = "Go"
	var framework = "Gin"
	var version string;
	jobTitle := "Backend Developer" // tanda : sebelah kiri dari = artinya deklarasi dan inisialisasi sekaligus

	var (
		firstName = "John"
		lastName = "Doe"
	)
	

	version = "1.20"
	jobTitle = "Fullstack Developer"

	fmt.Println("Programming Language:", language)
	fmt.Println("Web Framework:", framework)
	fmt.Println("Version:", version)
	fmt.Println("Job Title:", jobTitle)
	fmt.Println("Full Name:", firstName, lastName)
}