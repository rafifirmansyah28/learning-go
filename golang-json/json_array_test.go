package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName: "Rafi",
		LastName:  "Nanda",
		Age:       24,
		Married:   false,
		Hobbies:   []string{"Gaming", "Coding", "Reading"}, // Slice is representation of JSON Array
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Rafi","LastName":"Nanda","Age":24,"Married":false,"Hobbies":["Gaming","Coding","Reading"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Rafi",
		Addresses: []Address{
			{Street: "Jalan Mawar No 1", Country: "Indonesia", PostalCode: "12345"},
			{Street: "Jalan Melati No 2", Country: "Malaysia", PostalCode: "67890"},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Rafi","Addresses":[{"Street":"Jalan Mawar No 1","Country":"Indonesia","PostalCode":"12345"},{"Street":"Jalan Melati No 2","Country":"Malaysia","PostalCode":"67890"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Mawar No 1",
			Country:    "Indonesia",
			PostalCode: "12345",
		},
		{
			Street:     "Jalan Melati No 2",
			Country:    "Malaysia",
			PostalCode: "67890",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
