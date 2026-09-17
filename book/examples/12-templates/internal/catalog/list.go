package catalog

// Books creates the ordered initial catalog from the validated file title.
func Books(title string) ([]Book, bool) {
	books := []Book{
		{ID: "go-shop", Title: title, PriceMinor: 249900, Currency: "KZT"},
		{ID: "go-tests", Title: "Тесты на Go", PriceMinor: 150000, Currency: "KZT"},
	}
	for i := range books {
		if !books[i].Publish() {
			return nil, false
		}
	}
	return books, true
}
