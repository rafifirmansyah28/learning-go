package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Rafi"
	middleName = "Fadhlurrohman"
	lastName = "Sulaiman"

	return firstName, middleName, lastName
}
func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}