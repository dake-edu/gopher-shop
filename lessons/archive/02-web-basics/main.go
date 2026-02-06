package main

import (
	"fmt"
	"log"
	"net/http"
)

// Visual Anchor: The Container
type Book struct {
	Title  string
	Author string
	Price  float64
}

func main() {
	// The Shelf (Data)
	library := []Book{
		{Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Price: 35.99},
		{Title: "Learning Go", Author: "Jon Bodner", Price: 29.99},
	}

	// Visual Anchor: The Waiter (Handler)
	// When a customer (browser) asks for "/" (root), the waiter replies with this function.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("New Order Received!") // Log to console

		// The Food (HTML Response)
		// We write directly to the ResponseWriter 'w'.
		fmt.Fprintf(w, "<h1>Welcome to the Gopher Shop</h1>")
		fmt.Fprintf(w, "<ul>")
		for _, book := range library {
			fmt.Fprintf(w, "<li><b>%s</b> by %s ($%.2f)</li>", book.Title, book.Author, book.Price)
		}
		fmt.Fprintf(w, "</ul>")
	})

	// Open the Cafe
	fmt.Println("Server listing on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
