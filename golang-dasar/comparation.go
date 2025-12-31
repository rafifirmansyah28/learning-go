package main

import "fmt"

func main() {
	var nameA = "Rafi"
	var nameB = "Rafi"

	var result bool = nameA == nameB
	var result2 bool = nameA != nameB

	fmt.Println("result", result)
	fmt.Println("result2", result2)

	const thisYear = 2025
	const nextYear = 2026

	var result3 bool = thisYear < nextYear
	fmt.Println("result3", result3)
}