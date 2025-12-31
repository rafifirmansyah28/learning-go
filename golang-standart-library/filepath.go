package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("/home/rafi/docs/file.txt")) // Output: /home/rafi/docs
	fmt.Println(filepath.Base("/home/rafi/docs/file.txt")) // Output: file.txt
	fmt.Println(filepath.Ext("/home/rafi/docs/file.txt")) // Output: .txt
	fmt.Println(filepath.IsAbs("/home/rafi/docs/file.txt")) // Output: true
	fmt.Println(filepath.IsLocal("/home/rafi/docs/file.txt")) // Output: true
	fmt.Println(filepath.Join("home", "rafi", "docs")) // Output: home/rafi/docs
}