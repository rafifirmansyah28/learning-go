package main

import (
	"golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Rafi")
	fmt.Println(result)

	fmt.Println(helper.Application)
}