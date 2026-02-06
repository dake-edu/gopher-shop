package main

import (
	"fmt"
	"net/http"
)

// Lesson 1: The Server
// Goal: Listen for connections and talk back.

func main() {
	// 1. The Handler (The Clerk)
	// When someone visits "/", run this function.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello! Welcome to The Gopher Shop (v1)")
	})

	// 2. The Listener (The Open Sign)
	// Listen on port 8080.
	fmt.Println("🚀 Shop is open at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
