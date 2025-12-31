package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`r([a-z]+)i`)

	fmt.Print(regex.MatchString("raffi"))
	fmt.Print(regex.MatchString("rifki"))
	fmt.Print(regex.MatchString("rofi"))

	fmt.Println(regex.FindAllString("raffi rifi rifa rofi rafa", -1))
}