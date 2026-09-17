package main

import "fmt"

func main() {
	const currency = "KZT"
	const minorPerUnit = 100
	var priceMinor int64 = 249900
	title := "Go: от первой строки до интернет-магазина"
	published := false

	fmt.Println(title)
	fmt.Printf("Цена: %d.%02d %s\n", priceMinor/minorPerUnit, priceMinor%minorPerUnit, currency)
	fmt.Println("Опубликована:", published)
	published = true
	fmt.Println("Опубликована:", published)
}
