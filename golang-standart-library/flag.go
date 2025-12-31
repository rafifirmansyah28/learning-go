package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define command-line flags -> name, default value, description
	username := flag.String("username", "guest", "your username")
	password := flag.String("password", "secret", "your password")
	host := flag.String("host", "localhost", "database host")
	port := flag.Int("port", 5432, "database port")

	// Parse the command-line flags
	flag.Parse()

	fmt.Println("Username:", *username) // Dereference pointer to get the value
	fmt.Println("Password:", *password)
	fmt.Println("Host:", *host)
	fmt.Println("Port:", *port)
}