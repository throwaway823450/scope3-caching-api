package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("API_KEY")
    if apiKey == "" {
        log.Fatal("Unable to load api key from env API_KEY")
    }
	apiPath := os.Getenv("API_PATH")
    if apiPath == "" {
        log.Fatal("Unable to load api path from env API_PATH")
    }
	fmt.Println("Calling API path: ", apiPath)
}
