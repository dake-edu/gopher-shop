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
	// 1. The Shelf (Slice of Books)
	library := []Book{
		{Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Price: 35.99},
		{Title: "Learning Go", Author: "Jon Bodner", Price: 29.99},
		{Title: "Cloud Native Go", Author: "Matthew A. Titmus", Price: 34.99},
		{Title: "Go in Action", Author: "William Kennedy", Price: 24.99},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>The Gopher Shop Library</h1>")
		fmt.Fprintf(w, "<table border='1' cellpadding='10'>")
		fmt.Fprintf(w, "<tr><th>Title</th><th>Author</th><th>Price</th></tr>")

		// 2. The Conveyor Belt (For Range Loop)
		// We process each book one by one.
		for _, book := range library {
			fmt.Fprintf(w, "<tr>")
			fmt.Fprintf(w, "<td>%s</td>", book.Title)
			fmt.Fprintf(w, "<td>%s</td>", book.Author)
			fmt.Fprintf(w, "<td>$%.2f</td>", book.Price)
			fmt.Fprintf(w, "</tr>")
		}

		fmt.Fprintf(w, "</table>")
	})

	fmt.Println("Server running on http://localhost:8081 ...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
