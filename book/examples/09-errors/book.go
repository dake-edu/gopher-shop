package main

type Book struct {
	ID         string
	Title      string
	PriceMinor int64
	Currency   string
	Published  bool
}

// PriceAfterDiscount вычисляет цену, не изменяя книгу.
func (b Book) PriceAfterDiscount(percent int64) (int64, bool) {
	return priceAfterDiscount(b.PriceMinor, percent)
}

// Publish меняет состояние книги только после успешной проверки.
func (b *Book) Publish() bool {
	if b == nil || b.ID == "" || b.Currency != "KZT" {
		return false
	}
	title, ok := normalizeTitle(b.Title)
	if !ok {
		return false
	}
	_, ok = priceAfterDiscount(b.PriceMinor, 0)
	if !ok {
		return false
	}
	b.Title = title
	b.Published = true
	return true
}
