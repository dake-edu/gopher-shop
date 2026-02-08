package store

import (
	"database/sql"
	"fmt"

	"github.com/dake-edu/gopher-shop/internal/models"
	_ "github.com/lib/pq" // Register Postgres driver
)

// PostgresBookStore implements BookRepository using PostgreSQL.
type PostgresBookStore struct {
	db *sql.DB
}

// NewPostgresBookStore creates a new store.
func NewPostgresBookStore(db *sql.DB) *PostgresBookStore {
	return &PostgresBookStore{db: db}
}

// All returns all books from the database.
func (s *PostgresBookStore) All() ([]models.Book, error) {
	query := `SELECT id, title, author, price FROM books`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Price); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		books = append(books, b)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}
	return books, nil
}

// GetByID returns a book by ID.
func (s *PostgresBookStore) GetByID(id int) (*models.Book, bool, error) {
	query := `SELECT id, title, author, price FROM books WHERE id = $1`
	row := s.db.QueryRow(query, id)

	var b models.Book
	err := row.Scan(&b.ID, &b.Title, &b.Author, &b.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("scan failed: %w", err)
	}
	return &b, true, nil
}

// Create adds a new book to the database.
func (s *PostgresBookStore) Create(book *models.Book) error {
	query := `
		INSERT INTO books (title, author, price) 
		VALUES ($1, $2, $3) 
		RETURNING id`

	return s.db.QueryRow(query, book.Title, book.Author, book.Price).Scan(&book.ID)
}
