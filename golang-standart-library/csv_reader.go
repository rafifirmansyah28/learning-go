package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "jhon,doe,30\n" + "jane,doe,25\n" + "alice,smith,28"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read() // result is []string (slice of strings)
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}