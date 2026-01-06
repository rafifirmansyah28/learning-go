package golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Rafi",
		MiddleName: "Fajar",
		LastName:   "Nanda",
		Age:        24,
		Married:    false,
	}

	encoder.Encode(customer)
}
