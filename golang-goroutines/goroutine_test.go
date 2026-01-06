package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	// RunHelloWorld() // run normally
	go RunHelloWorld() // run as goroutine

	fmt.Println("Ups")

	time.Sleep(1 * time.Second) // give time for goroutine to finish
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := range 1000000 {
		go DisplayNumber(i)
	}

	time.Sleep(100 * time.Second)
}
