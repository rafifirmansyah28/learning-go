package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Rafi", "Maulana", "Putra", "Pratama"}
	values := []int{10, 20, 30, 40, 50}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Rafi"))
	fmt.Println(slices.Index(names, "John"))
	fmt.Println(slices.Index(names, "Pratama"))
	fmt.Println(slices.Index(values, 30))

}