package service

import (
	"errors"
	"ims/models"
)

type ProductStorage interface {
	Save(p models.Product) error
	GetByID(id int) (*models.Product, error)
	Delete(id int) error
}

type MockProductStorage struct {
	prods map[int]models.Product
}

func NewMockStorage() *MockProductStorage {
	return &MockProductStorage{prods: make(map[int]models.Product)}
}

func (m *MockProductStorage) Save(p models.Product) error {
	if _, exists := m.prods[p.ID]; exists {
		return errors.New("product already exists")
	}
	m.prods[p.ID] = p
	return nil
}
func (m *MockProductStorage) GetByID(id int) (*models.Product, error) {
	if product, exists := m.prods[id]; exists {
		return &product, nil
	}
	return nil, errors.New("product not found")
}
func (m *MockProductStorage) Delete(id int) error {
	if _, exists := m.prods[id]; exists {
		delete(m.prods, id)
		return nil
	}
	return errors.New("product not found")
}
