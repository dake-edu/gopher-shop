package main

import "fmt"

func main() {
	book := Book{
		ID:         "go-shop",
		Title:      "  Go: от первой строки до книжного магазина  ",
		PriceMinor: 249900,
		Currency:   "KZT",
	}
	fmt.Println("До публикации:", book.Published)
	if !book.Publish() {
		fmt.Println("Книга не готова к публикации")
		return
	}
	fmt.Println("После публикации:", book.Published)
	price, ok := book.PriceAfterDiscount(10)
	if !ok {
		fmt.Println("Скидка не подходит")
		return
	}
	fmt.Println(book.Title)
	fmt.Println("К оплате:", price, "минимальных единиц", book.Currency)
}
