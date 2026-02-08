package store

import "github.com/dake-edu/gopher-shop/internal/models"

// ------------------------------------------------------------------------------------------------
// ⚓ VISUAL ANCHOR: THE REPOSITORY INTERFACE
// ------------------------------------------------------------------------------------------------
//
//      [ Application ]  --->  ( Interface )  --->  [ Implementation ]
//
//      The App doesn't care HOW data is stored (Memory, Postgres, Redis).
//      It only cares that the store can "GetAll" and "GetByID".
//
// ------------------------------------------------------------------------------------------------

// BookRepository defines the behavior for book storage.
type BookRepository interface {
	// All returns a list of all books.
	// Note: We use All() instead of GetAll() following the Effective Go naming conventions.
	All() ([]models.Book, error)
	GetByID(id int) (*models.Book, bool, error)
	Create(book *models.Book) error
}
