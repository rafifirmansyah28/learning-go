package main

import (
	"golang-dasar/database"
	_ "golang-dasar/internal" // underscore (blank identifier) digunakan untuk mengimpor package hanya untuk menjalankan fungsi init saja
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}