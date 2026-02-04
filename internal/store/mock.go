package store

import "github.com/dake-edu/gopher-shop/internal/models"

// MockBookStore is a mock implementation of BookRepository for testing.
type MockBookStore struct {
	GetAllFunc  func() ([]models.Book, error)
	GetByIDFunc func(id int) (*models.Book, bool, error)
	CreateFunc  func(book *models.Book) error
}

func (m *MockBookStore) GetAll() ([]models.Book, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return nil, nil
}

func (m *MockBookStore) GetByID(id int) (*models.Book, bool, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(id)
	}
	return nil, false, nil
}

func (m *MockBookStore) Create(book *models.Book) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(book)
	}
	return nil
}
