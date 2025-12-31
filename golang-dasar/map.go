package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Rafi"
	// person["address"] = "Indonesia"

	person := map[string]string{
		"name":    "Rafi",
		"address": "Indonesia",
	}

	fmt.Println("Name :", person["name"])
	fmt.Println("Address :", person["address"])
	fmt.Println(person)
	fmt.Println(person["job_title"])

	book := make(map[string]string)
	book["title"] = "Learning Go"
	book["author"] = "Rafi"
	book["ups"] = "Buku ini bagus sekali"
	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
