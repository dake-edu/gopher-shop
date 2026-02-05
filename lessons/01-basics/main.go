package main

import "fmt"

// Visual Anchor: The Container
// A Struct is like a container with specific labels.
// Here, our container is labeled "Book" and holds a Title, Author, and Price.
type Book struct {
	Title  string
	Author string
	Price  float64
}

func main() {
	fmt.Println("Welcome to the Gopher Shop - Junior Path!")

	// 1. Creating individual containers (Structs)
	book1 := Book{
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan",
		Price:  35.99,
	}

	book2 := Book{
		Title:  "Learning Go",
		Author: "Jon Bodner",
		Price:  29.99,
	}

	// 2. Placing containers on a Shelf (Slice)
	// A Slice is an ordered collection, just like a shelf where you place books one by one.
	// We call our shelf "library".
	library := []Book{book1, book2}

	// 3. Viewing the Shelf
	// We walk along the shelf and look at each book.
	fmt.Println("\n--- My Book Shelf ---")
	for i, book := range library {
		fmt.Printf("%d. %s by %s ($%.2f)\n", i+1, book.Title, book.Author, book.Price)
	}
}
