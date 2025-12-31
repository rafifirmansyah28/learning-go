package main

import (
	"fmt"
	"strconv"
)

func main() {
	// convert string to boolean
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Result:", result)
	}

	// convert string to integer
	resultInt, err := strconv.Atoi("100")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("resultInt:", resultInt)
	}

	// convert integer to binary
	binary := strconv.FormatInt(999, 2)
	fmt.Println("binary:", binary)

	// convert integer to string
	var stringInt = strconv.Itoa(100)
	fmt.Println("stringInt:", stringInt)
}