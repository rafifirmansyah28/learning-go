package main

import "fmt"

func main() {
	names := [...]string{"Rafi", "Fauzi", "Aditya", "Putra", "Saputra", "Bagus"}

	slice1 := names[1:4]
	fmt.Println(slice1)

	slice2 := names[0:3]
	fmt.Println(slice2)

	var slice3 []string = names[2:5]
	fmt.Println(slice3)

	var slice4 []string = names[1:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Rafi"
	newSlice[1] = "Fauzi"
	// newSlice[2] = "Aditya" // error, must use append to add data if exceed length
	fmt.Println(newSlice)
	fmt.Println("Length =", len(newSlice))
	fmt.Println("Capacity =", cap(newSlice))

	newSlice2 := append(newSlice, "Aditya")
	fmt.Println(newSlice2)
	fmt.Println("Length =", len(newSlice2))
	fmt.Println("Capacity =", cap(newSlice2))

	months := []string{"Januari", "Februari", "Maret", "April", "Mei"}
	fmt.Println(months)

	monthsSlice1 := months[2:3]
	fmt.Println(monthsSlice1)
	fmt.Println("Length =", len(monthsSlice1))
	fmt.Println("Capacity =", cap(monthsSlice1))
	monthsSlice1 = append(monthsSlice1, "Maret Baru")
	fmt.Println(monthsSlice1)
	fmt.Println("Length =", len(monthsSlice1))
	fmt.Println("Capacity =", cap(monthsSlice1))
	monthsSlice2 := append(monthsSlice1, "April Baru")
	fmt.Println(monthsSlice1)
	fmt.Println(monthsSlice2)
	fmt.Println(months)

	// Copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	thisIsArray := [...]int{1, 2, 3, 4, 5}
	thisIsSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(thisIsArray)
	fmt.Println(thisIsSlice)
}