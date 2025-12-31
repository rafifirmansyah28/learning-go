package main

import "fmt"

func main() {
	var (
		a = 10
		b = 20
		d = 50
		e = 40
	)
	var z = a + b - d * e
	fmt.Println(z)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	i += 5 // i = i + 5
	fmt.Println(i)

	var j = 1
	j++ // j = j + 1
	fmt.Println(j)

	j--
	fmt.Println(j)
}