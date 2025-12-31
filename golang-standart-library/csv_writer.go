package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"jhon", "doe", "30"})
	_ = writer.Write([]string{"jane", "doe", "25"})
	_ = writer.Write([]string{"alice", "smith", "28"})

	writer.Flush()
}