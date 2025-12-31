package main

import (
	"errors"
	"fmt"
)

// Define custom error variables
var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

// Simulate a function that returns different errors based on input
func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "rafi" {
		return NotFoundError
	}

	// Simulate successful retrieval
	return nil
}

func main() {
	err := GetById("")

	if err != nil {
		// erros.Is to check for specific error types
		if errors.Is(err, ValidationError) { 
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	}
}