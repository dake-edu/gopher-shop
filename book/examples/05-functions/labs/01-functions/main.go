package main

import "fmt"

func announce(title string) {
	fmt.Println("Книга:", title)
}
func total(priceMinor int64, quantity int64) int64 {
	return priceMinor * quantity
}
func changeQuantity(quantity int64) {
	quantity = 10
	fmt.Println("Внутри:", quantity)
}
func main() {
	announce("Go")
	var quantity int64 = 2
	amount := total(150, quantity)
	fmt.Println("Сумма:", amount)
	changeQuantity(quantity)
	fmt.Println("Снаружи:", quantity)
}
