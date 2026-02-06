package main

import (
	"fmt"
)

// Lesson 5: Maps (The Catalog Index)
// Goal: Learn how to store Key-Value pairs.
// Visual Anchor: "The Coat Check" (Give Ticket -> Get Coat).

func main() {
	// 1. Creating the Map (The Coat Check Room)
	// Why map[string]float64?
	// - string: The Key (The "Ticket"), e.g., "The Go Gopher"
	// - float64: The Value (The "Coat"), e.g., 25.00
	prices := map[string]float64{
		"The Go Gopher":     25.00,
		"Concurrency in Go": 30.00,
		"Black Hat Go":      35.50,
	}

	// 2. Lookup (Giving the Ticket)
	bookName := "The Go Gopher"
	price, exists := prices[bookName]

	// Why "exists"?
	// Go maps tell you TWO things:
	// 1. The value (price)
	// 2. Whether it was found (exists) - Essential for safety!
	if exists {
		fmt.Printf("found '%s': $%.2f\n", bookName, price)
	} else {
		fmt.Printf("Sorry, we don't carry '%s'\n", bookName)
	}

	// 3. Adding/Updating (Checking in a new Coat)
	prices["Blueprints"] = 40.00
	fmt.Println("Added 'Blueprints'. Current Catalog size:", len(prices))

	// 4. Deleting
	delete(prices, "Black Hat Go")
}
