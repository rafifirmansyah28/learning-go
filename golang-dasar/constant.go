package main

import "fmt"

func main() {
	const test = "This is constant"
	test = "Trying to change constant" // will cause error
	const (
		a = 1
		b = 2
		c = 3
	) // create multiple constant

	fmt.Println(test)
}