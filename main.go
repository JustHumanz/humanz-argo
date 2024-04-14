package main

import (
	"fmt"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	version := getEnv("VERSION", "v.0.0")
	// Register the handler function for the root ("/") route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write the response to the client
		fmt.Fprint(w, "Hello, Humanz!")
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		// Write the response to the client
		fmt.Fprint(w, "App version ", version)
	})

	if err := http.ListenAndServe(":2525", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
