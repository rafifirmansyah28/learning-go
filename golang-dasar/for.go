package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	fmt.Println("Selesai")

	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Perulangan ke ", counter2)
	}

	// For range -> array/slice/map

	names := []string{"Rafi", "Adit", "Budi", "Caca"}

	for i := 0; i < len(names); i++ {
		fmt.Println("Index", i, "=", names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println("Name", name)
	}
}