package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	title, err := loadTitle("title.txt")
	if err != nil {
		return err
	}
	book := Book{ID: "go-shop", Title: title, PriceMinor: 249900, Currency: "KZT"}
	if !book.Publish() {
		return errors.New("карточка книги не прошла проверку")
	}
	fmt.Println(book.Title)
	fmt.Println("Опубликована:", book.Published)
	return nil
}
