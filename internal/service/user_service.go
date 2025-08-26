// Package service provides business logic for user-related operations.
package service

import (
	"github.com/matheuszuge/golang/internal/domain"
	"github.com/matheuszuge/golang/internal/repository"
)

type UserService interface {
	ListUsers() ([]domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) ListUsers() ([]domain.User, error) {
	return s.repo.List()
}
