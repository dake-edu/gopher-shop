package service

import (
	"errors"
	"github.com/dake-edu/gopher-shop/internal/models"
	"github.com/dake-edu/gopher-shop/internal/store"
	"testing"
)

func TestInvalidBookNeverReachesStore(t *testing.T) {
	writes := 0
	repo := &store.MockBookStore{CreateFunc: func(*models.Book) error { writes++; return nil }}
	s := NewBookService(repo)
	for _, b := range []*models.Book{nil, {Title: "Go", Price: 1}, {Title: "Go", Author: "Author", Price: 0}} {
		if s.CreateBook(b) == nil {
			t.Fatal("invalid book accepted")
		}
	}
	if writes != 0 {
		t.Fatalf("invalid writes: %d", writes)
	}
	if err := s.CreateBook(&models.Book{Title: "Go", Author: "Author", Price: 1}); err != nil {
		t.Fatal(err)
	}
	if writes != 1 {
		t.Fatalf("valid writes: %d", writes)
	}
}
func TestMissingBookHasRecognizableCause(t *testing.T) {
	_, err := NewBookService(&store.MockBookStore{}).GetBook(999)
	if !errors.Is(err, ErrBookNotFound) {
		t.Fatalf("missing cause: %v", err)
	}
}
