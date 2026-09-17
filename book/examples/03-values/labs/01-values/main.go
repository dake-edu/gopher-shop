package main

import "fmt"

func main() {
	var title string
	var priceMinor int64
	var published bool
	fmt.Printf("title=[%s], price=%d, published=%t\n", title, priceMinor, published)
	title = "Go"
	otherTitle := title
	title = "Go shop"
	fmt.Println(title)
	fmt.Println(otherTitle)
	quantity := 2
	priceMinor = 150
	total := int64(quantity) * priceMinor
	fmt.Printf("%T %T %d\n", quantity, priceMinor, total)
	const percent = 10
	var discount int64 = percent
	fmt.Println(discount)
}
