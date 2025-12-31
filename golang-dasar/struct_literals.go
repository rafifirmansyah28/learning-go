package main

import "fmt"

type Customer struct {
	Name, Address    string
	Age     int
}

// method adalah function yang ada di dalam struct. Bisa dipanggil menggunakan receiver dari struct tersebut
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	joko := Customer{
		Name:   "Joko",
		Address: "Indonesia",
		Age:    30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 25}
	fmt.Println(budi)

	budi.sayHello("Joni")
}