package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"scope3/caching-api/api"
	"scope3/caching-api/measurement"
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
	fmt.Println("Using API path: ", apiPath)
	measurementClient := measurement.NewClient(apiPath, apiKey)
	cachingClient := measurement.NewCachingClient(measurementClient)

	router := api.NewRouter(cachingClient)

	log.Fatal(http.ListenAndServe(":8080", router))
}
