package main

import (
	"fmt"
	"log"

	"github.com/o5h/config"
)

func main() {
	// Load configuration from a YAML file
	err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	// Access configuration values
	fmt.Println("Server Port:", config.Get("server.port", ""))
	fmt.Println("Database Host:", config.Get("database.host", ""))
}
