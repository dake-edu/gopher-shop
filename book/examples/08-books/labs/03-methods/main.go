package main

import "fmt"

type Book struct {
	Title     string
	Published bool
}

func (book Book) RenameCopy(title string) {
	book.Title = title
}
func (book *Book) Publish() bool {
	if book == nil {
		return false
	}
	book.Published = true
	return true
}
func main() {
	book := Book{Title: "Go"}
	book.RenameCopy("SQL")
	fmt.Println(book.Title)
	fmt.Println(book.Publish(), book.Published)
	var absent *Book
	fmt.Println(absent.Publish())
	books := []Book{{Title: "Go"}, {Title: "SQL"}}
	for _, copy := range books {
		copy.Publish()
	}
	fmt.Println(books[0].Published, books[1].Published)
	for i := range books {
		books[i].Publish()
	}
	fmt.Println(books[0].Published, books[1].Published)
}
