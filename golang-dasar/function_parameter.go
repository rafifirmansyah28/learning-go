package main

import "fmt"

func sayHelloTo(name string) {
	fmt.Println("Hello", name)
}

func main() {
	sayHelloTo("Rafi")
	sayHelloTo("Budi")
	sayHelloTo("Ani")
}