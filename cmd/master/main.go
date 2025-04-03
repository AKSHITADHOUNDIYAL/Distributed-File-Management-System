package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Distributed File System - Master Node")

	// Example: Start a simple HTTP server
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Master Node is Running")
	})

	port := ":8080"
	fmt.Println("Starting server on", port)
	http.ListenAndServe(port, nil)
}
