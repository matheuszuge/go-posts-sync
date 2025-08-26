// Package service provides business logic for user-related operations.
package service

import (
	"github.com/matheuszuge/golang/internal/domain"
	"github.com/matheuszuge/golang/internal/repository"
)

type PostService interface {
	ListPosts() ([]domain.Post, error)
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(r repository.PostRepository) PostService {
	return &postService{repo: r}
}

func (s *postService) ListPosts() ([]domain.Post, error) {
	return s.repo.List()
}
