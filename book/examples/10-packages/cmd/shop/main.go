package main

import (
	"fmt"
	"os"

	"github.com/dake-edu/gopher-shop/book/examples/10-packages/internal/catalog"
	"github.com/dake-edu/gopher-shop/book/examples/10-packages/internal/titlefile"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	title, err := titlefile.Load("title.txt")
	if err != nil {
		return err
	}
	book := catalog.Book{ID: "go-shop", Title: title, PriceMinor: 249900, Currency: "KZT"}
	if !book.Publish() {
		return fmt.Errorf("карточка книги не прошла проверку")
	}
	fmt.Println(book.Title)
	fmt.Println("Опубликована:", book.Published)
	return nil
}
