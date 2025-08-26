// Package repository provides user repository implementations.
package repository

import "github.com/matheuszuge/golang/internal/domain"

type UserRepository interface {
	List() ([]domain.User, error)
}

type InMemoryUserRepo struct {
	data []domain.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		data: []domain.User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}},
	}
}

func (r *InMemoryUserRepo) List() ([]domain.User, error) {
	return r.data, nil
}
