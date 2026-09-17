package main

import "fmt"

func main() {
	var uninitialized map[string]int
	fmt.Println(uninitialized["go"])
	quantities := make(map[string]int)
	quantities["go"] = 0
	value, exists := quantities["go"]
	missing, found := quantities["sql"]
	fmt.Println(value, exists, missing, found)
	quantities["go"] = 2
	same := quantities
	same["go"] = 3
	fmt.Println(quantities["go"])
	delete(quantities, "go")
	_, exists = quantities["go"]
	fmt.Println(exists, len(quantities))
	quantities["sql"] = 1
	clear(quantities)
	fmt.Println(len(quantities), quantities == nil)
}
