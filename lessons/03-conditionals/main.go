package main

import (
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Title  string
	Author string
	Price  float64
}

func main() {
	// 1. The Data (Shelf)
	library := []Book{
		{Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Price: 35.99},
		{Title: "Learning Go", Author: "Jon Bodner", Price: 29.99},
	}

	// 2. The Logic Switch (Boolean)
	// Try changing this to false!
	isOpen := true

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 3. The Fork in the Road (If/Else)
		if isOpen {
			// Path A: The Shop is Open
			fmt.Fprintf(w, "<h1>Welcome to the Gopher Shop!</h1>")
			fmt.Fprintf(w, "<ul>")
			for _, book := range library {
				fmt.Fprintf(w, "<li>%s ($%.2f)</li>", book.Title, book.Price)
			}
			fmt.Fprintf(w, "</ul>")
		} else {
			// Path B: The Shop is Closed
			fmt.Fprintf(w, "<h1>Sorry, we are closed.</h1>")
			fmt.Fprintf(w, "<p>The Gopher is taking a nap. Come back later!</p>")
		}
	})

	fmt.Println("Server running on http://localhost:8081 ...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
