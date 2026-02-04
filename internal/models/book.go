package models

import (
	"fmt"
)

// ------------------------------------------------------------------------------------------------
// ⚓ VISUAL ANCHOR: THE BOOK ENTITY
// ------------------------------------------------------------------------------------------------
//
//      +------------------------------------------+
//      |                  Book                    |
//      +==========================================+
//      | ID      (int)     : Unique Identifier    |
//      | Title   (string)  : Name of the book     |
//      | Author  (string)  : Writer's name        |
//      | Price   (float64) : Cost in USD          |
//      +------------------------------------------+
//
// ------------------------------------------------------------------------------------------------

// Book represents a book in our store.
type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// Validate checks if the book has valid data.
func (b *Book) Validate() error {
	if b.Title == "" {
		return fmt.Errorf("title is required")
	}
	if b.Author == "" {
		return fmt.Errorf("author is required")
	}
	if b.Price <= 0 {
		return fmt.Errorf("price must be positive")
	}
	return nil
}
