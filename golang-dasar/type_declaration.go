package main

import "fmt"

func main() {
	type NoKTP string
	
	var ktpRafi NoKTP = "3573041506000002"
	fmt.Println(ktpRafi)

	var contoh string = "222222";
	var contohKtp = NoKTP(contoh)

	fmt.Println(contohKtp)
}