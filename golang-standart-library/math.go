package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.35)) // pembulatan ke atas
	fmt.Println(math.Floor(1.35)) // pembulatan ke bawah
	fmt.Println(math.Round(1.50)) // pembulatan ke terdekat
	fmt.Println(math.Max(10, 20)) // nilai maksimum
	fmt.Println(math.Min(10, 20)) // nilai minimum
}