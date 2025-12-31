package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("/home/rafi/docs/file.txt")) // Output: /home/rafi/docs
	fmt.Println(path.Base("/home/rafi/docs/file.txt")) // Output: file.txt
	fmt.Println(path.Ext("/home/rafi/docs/file.txt")) // Output: .txt
	fmt.Println(path.Join("home", "rafi", "docs")) // Output: home/rafi/docs
}