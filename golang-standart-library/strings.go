package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Rafi Maulana", "Rafi")) // search substring
	fmt.Println(strings.Split("Rafi,Maulana", ",")) // split string by separator
	fmt.Println(strings.ToLower("Rafi Maulana")) // convert to lowercase
	fmt.Println(strings.ToUpper("Rafi Maulana")) // convert to uppercase
	fmt.Println(strings.Trim("   Rafi Maulana   ", " ")) // trim spaces from both ends
	fmt.Println(strings.ReplaceAll("Rafi Maulana", "a", "o")) // replace all occurrences of a substring
}