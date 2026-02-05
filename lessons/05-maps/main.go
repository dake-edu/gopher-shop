package main

import "fmt"

func main() {
	// ---------------------------------------------------------
	// ⚓ CONCEPT: MAPS (The Coat Check)
	// ---------------------------------------------------------
	// A Map connects a "Key" (Ticket) to a "Value" (Coat).
	// Unlike Arrays, they are not numbered 0, 1, 2.
	// You look things up by their Label.
	// ---------------------------------------------------------

	// 1. Create a Map (The Empty Rack)
	//    Key: string (Product Name)
	//    Value: float64 (Price)
	menu := make(map[string]float64)

	// 2. Add Items
	menu["Coffee"] = 2.50
	menu["Tea"] = 1.75
	menu["Cookie"] = 1.00

	// 3. Look up an item
	fmt.Println("Price of Coffee:", menu["Coffee"])

	// 4. Check if item exists (The "Comma OK" idiom)
	//    price: The value (0 if missing)
	//    exists: True/False
	price, exists := menu["Pizza"]
	if !exists {
		fmt.Println("Sorry, we don't serve Pizza.")
	} else {
		fmt.Println("Pizza costs:", price)
	}

	// 5. Delete an item
	delete(menu, "Cookie")
	fmt.Println("Menu size after deleting cookie:", len(menu))

	// 6. Loop (Random Order!)
	fmt.Println("\n--- Full Menu ---")
	for name, price := range menu {
		fmt.Printf("%s: $%.2f\n", name, price)
	}
}
