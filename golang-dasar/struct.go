package main

import "fmt"

// struct adalah kumpulan dari beberapa tipe data yang bisa berbeda atau bahasa simplenya adalah template data/prototype data
type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var rafi Customer
	rafi.Name = "Rafi"
	rafi.Address = "Indonesia"
	rafi.Age = 20

	fmt.Println(rafi)
	fmt.Println(rafi.Name)
}