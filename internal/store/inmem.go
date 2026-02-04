package store

import "github.com/dake-edu/gopher-shop/internal/models"

// InMemoryBookStore implements BookRepository using a slice.
type InMemoryBookStore struct {
	books []models.Book
}

// NewInMemoryBookStore creates a new store with sample data.
func NewInMemoryBookStore() *InMemoryBookStore {
	return &InMemoryBookStore{
		books: []models.Book{
			{ID: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Price: 35.99},
			{ID: 2, Title: "Clean Architecture", Author: "Robert C. Martin", Price: 28.50},
			{ID: 3, Title: "Domain-Driven Design", Author: "Eric Evans", Price: 42.00},
		},
	}
}

// GetAll returns all books.
func (s *InMemoryBookStore) GetAll() ([]models.Book, error) {
	return s.books, nil
}

// GetByID returns a book by ID.
func (s *InMemoryBookStore) GetByID(id int) (*models.Book, bool, error) {
	for _, book := range s.books {
		if book.ID == id {
			return &book, true, nil
		}
	}
	return nil, false, nil
}

// Create adds a new book to the in-memory store.
func (s *InMemoryBookStore) Create(book *models.Book) error {
	book.ID = len(s.books) + 1
	s.books = append(s.books, *book)
	return nil
}

// Kept for backward compatibility during refactor if needed, or remove.
// We will simply replace the usage in main.go, so I'll remove the global var to force the refactor.
