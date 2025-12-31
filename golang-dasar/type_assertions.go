package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var result any = random()
	// var resultString string = result.(string) // type assertion
	// fmt.Println(resultString)

	// var resultInt int = result.(int) // akan error karena result bukan int
	// fmt.Println(resultInt)

	switch value := result.(type) { // type assertion dengan switch untuk menghindari error
		case string:
			fmt.Println("String", value)
		case int:
			fmt.Println("Integer", value)
		default:
			fmt.Println("Unknown")
	}

}