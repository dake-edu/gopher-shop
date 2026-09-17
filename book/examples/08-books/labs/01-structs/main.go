package main

import "fmt"

type Book struct {
	Title      string
	PriceMinor int64
	Published  bool
}

func main() {
	original := Book{Title: "Go", PriceMinor: 150}
	copied := original
	copied.Title = "SQL"
	fmt.Println(original.Title, copied.Title, original.Published)
	var empty Book
	fmt.Printf("[%s] %d %t\n", empty.Title, empty.PriceMinor, empty.Published)
}
