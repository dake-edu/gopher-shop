package service

import (
	"fmt"
	"log/slog"

	"github.com/dake-edu/gopher-shop/internal/models"
	"github.com/dake-edu/gopher-shop/internal/store"
)

// BookService is the "Brain" of the operation.
// It handles business logic and rules before talking to the "Warehouse" (Store).
type BookService struct {
	repo store.BookRepository
}

// NewBookService creates a new "Brain" (Service) and gives it a "Warehouse" (Repo) to usage.
func NewBookService(repo store.BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

// ListBooks fetches all books and applies business logic (if any).
// For now, it just asks the warehouse, but later we could filter out "Draft" books here.
func (s *BookService) ListBooks(category string) ([]models.Book, error) {
	slog.Debug("listing books", "category", category)

	allBooks, err := s.repo.All()
	if err != nil {
		return nil, fmt.Errorf("service: could not fetch books: %w", err)
	}

	// 🧠 Business Logic: Filtering happens here (The Brain), not in the Handler (The Front Desk).
	// Although for huge databases, we might push this down to SQL (SELECT WHERE...).
	// But for our "Learning Mode", this is where logic lives.
	if category == "" {
		return allBooks, nil
	}

	var filtered []models.Book
	for _, b := range allBooks {
		if b.Category == category {
			filtered = append(filtered, b)
		}
	}
	return filtered, nil
}

// GetBook fetches a single book.
func (s *BookService) GetBook(id int) (models.Book, error) {
	// The Service Layer delegates to the Store Layer.
	// But it could also check cache first! (Future Enhancement)
	book, found, err := s.repo.GetByID(id)
	if err != nil {
		return models.Book{}, fmt.Errorf("service: db error: %w", err)
	}
	if !found {
		return models.Book{}, fmt.Errorf("service: book not found")
	}
	return *book, nil
}

// CreateBook adds a new book to the shop.
// 🧠 Business Rule: Logic Layer ensures data is valid before saving.
// E.g., "Price cannot be negative".
func (s *BookService) CreateBook(b *models.Book) error {
	slog.Debug("validating new book", "title", b.Title, "price", b.Price)

	// 1. Validation Rule
	if b.Price < 0 {
		return fmt.Errorf("service: price cannot be negative (got %0.2f)", b.Price)
	}
	if b.Title == "" {
		return fmt.Errorf("service: book title cannot be empty")
	}

	// 2. Delegate to Store
	if err := s.repo.Create(b); err != nil {
		return fmt.Errorf("service: could not create book: %w", err)
	}

	return nil
}
