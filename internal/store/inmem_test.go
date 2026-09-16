package store

import (
	"github.com/dake-edu/gopher-shop/internal/models"
	"testing"
)

func TestCreateAssignsUnusedID(t *testing.T) {
	s := NewInMemoryBookStore()
	before, err := s.All()
	if err != nil {
		t.Fatal(err)
	}
	book := models.Book{Title: "New", Author: "Author", Price: 1}
	if err := s.Create(&book); err != nil {
		t.Fatal(err)
	}
	for _, old := range before {
		if old.ID == book.ID {
			t.Fatalf("duplicate ID %d", book.ID)
		}
	}
	got, found, err := s.GetByID(book.ID)
	if err != nil || !found || got.Title != "New" {
		t.Fatalf("created book unavailable: %v, %t, %v", got, found, err)
	}
}
