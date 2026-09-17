package catalog

import "testing"

func TestPublish(t *testing.T) {
	book := Book{ID: "go", Title: "  Go  ", PriceMinor: 105, Currency: "KZT"}
	if !book.Publish() || !book.Published || book.Title != "Go" {
		t.Fatalf("публикация не сохранилась: %+v", book)
	}
	price, ok := book.PriceAfterDiscount(10)
	if !ok || price != 95 || book.PriceMinor != 105 {
		t.Fatal("скидка неверна или изменила исходную цену")
	}
}

func TestInvalidBookUnchanged(t *testing.T) {
	book := Book{ID: "go", Title: "  Go  ", PriceMinor: -1, Currency: "KZT"}
	before := book
	if book.Publish() || book != before {
		t.Fatal("невалидная книга опубликована или изменена")
	}
}

func TestNilBook(t *testing.T) {
	var book *Book
	if book.Publish() {
		t.Fatal("nil книга опубликована")
	}
}

func TestMissingBookID(t *testing.T) {
	book := Book{Title: "Go", PriceMinor: 100, Currency: "KZT"}
	if book.Publish() {
		t.Fatal("книга без ID опубликована")
	}
}

func TestUnsupportedCurrency(t *testing.T) {
	book := Book{ID: "go", Title: "Go", PriceMinor: 100, Currency: "XXX"}
	if book.Publish() {
		t.Fatal("неподдерживаемая валюта принята")
	}
}

func TestRangeValueDoesNotPublishOriginal(t *testing.T) {
	books := []Book{{ID: "go", Title: "Go", PriceMinor: 100, Currency: "KZT"}}
	for _, book := range books {
		book.Publish()
	}
	if books[0].Published {
		t.Fatal("изменение копии неожиданно изменило исходный Published")
	}
	for i := range books {
		books[i].Publish()
	}
	if !books[0].Published {
		t.Fatal("изменение элемента среза не сохранилось")
	}
}
