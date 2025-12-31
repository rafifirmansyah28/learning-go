package main

import "fmt"

func main() {
	var names [3]string // static array -> length has defined

	names[0] = "Rafi"
	names[1] = "Fauzi"
	names[2] = "Ahmad"
	
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{10, 20, 30}
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	var test = [...]int{1, 2, 3, 4, 5} // dynamic array -> length will adjust with the data
	fmt.Println("Length array test =", len(test))

	fmt.Println(test)
	test[0] = 11
	fmt.Println(test[0])
}