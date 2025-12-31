package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil // nil adalah nilai kosong untuk tipe data reference seperti map, slice, pointer, channel, dll
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("")

	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}

}