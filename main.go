package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Register the handler function for the root ("/") route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write the response to the client
		fmt.Fprint(w, "Hello, Humanz!")
	})

	if err := http.ListenAndServe(":2525", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
