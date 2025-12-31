package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	rafi := Man{"Rafi"}
	rafi.Married()

	fmt.Println(rafi)
}