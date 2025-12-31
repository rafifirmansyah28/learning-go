package helper

// Variabel dengan huruf kecil di awal menandakan bahwa variabel ini hanya bisa diakses dari package yang sama
var version = "1.0.0"

var Application = "Golang"

// Function dengan huruf kapital di awal menandakan bahwa function ini bisa diakses dari package lain
func sayGoodBye(name string) string {
	return "Good Bye " + name
}

// Huruf kapital di awal menandakan bahwa function ini bisa diakses dari package lain
func SayHello(name string) string {
	return "Hello " + name
}