package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() // logging akan dipanggil di akhir function runApplication walaupun ada return/error di tengah
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}