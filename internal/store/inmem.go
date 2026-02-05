package store

import (
	"errors"
	"sync"

	"github.com/dake-edu/gopher-shop/internal/models"
)

// InMemoryBookStore is a simple thread-safe store for the demo.
// It matches the "Repository Pattern" but keeps data in RAM.
type InMemoryBookStore struct {
	books  []models.Book
	nextID int
	mu     sync.Mutex
}

// NewInMemoryBookStore creates a store with some initial data.
func NewInMemoryBookStore() *InMemoryBookStore {
	return &InMemoryBookStore{
		books: []models.Book{
			{ID: 1, Title: "The Go Gopher", Author: "Rob Pike", Price: 25.00},
			{ID: 2, Title: "Concurrency in Go", Author: "Katherine Cox-Buday", Price: 30.00},
		},
		nextID: 3,
	}
}

// GetAll returns a copy of all books (to avoid race conditions).
func (s *InMemoryBookStore) GetAll() ([]models.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Return a copy, not the original slice
	result := make([]models.Book, len(s.books))
	copy(result, s.books)
	return result, nil
}

// Create adds a new book to the store.
func (s *InMemoryBookStore) Create(book *models.Book) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	book.ID = s.nextID
	s.nextID++
	s.books = append(s.books, *book)
	return nil
}

// GetByID returns a single book.
func (s *InMemoryBookStore) GetByID(id int) (models.Book, bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, b := range s.books {
		if b.ID == id {
			return b, true, nil
		}
	}
	return models.Book{}, false, errors.New("not found")
}
