package main

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover() // recover menangkap panic dan mengembalikan nilai panic tersebut
	fmt.Println("Error message:", message)
}

func runApp(error bool) {
	defer endApp() // endApp akan dipanggil di akhir function runApp walaupun ada panic di tengah

	if error {
		panic("Ups Error")
	}
}

func main() {
	runApp(true)
}