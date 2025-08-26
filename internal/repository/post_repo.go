// Package repository provides user repository implementations.
package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/matheuszuge/golang/internal/domain"
)

type PostRepository interface {
	List() ([]domain.Post, error)
}

type InMemoryPostRepo struct {
	data []domain.Post
}

func NewInMemoryPostRepo(data []domain.Post) *InMemoryPostRepo {
	return &InMemoryPostRepo{data: data}
}

func (r *InMemoryPostRepo) List() ([]domain.Post, error) {
	return r.data, nil
}

type APIPostRepo struct {
	BaseURL string
	Client  *http.Client
}

func NewAPIPostRepo(baseURL string, client *http.Client) *APIPostRepo {
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}

	return &APIPostRepo{BaseURL: baseURL, Client: client}
}

func (r *APIPostRepo) List() ([]domain.Post, error) {
	req, err := http.NewRequest(http.MethodGet, r.BaseURL+"/posts", nil)
	if err != nil {
		return nil, err
	}
	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote api returned %d", resp.StatusCode)
	}

	var posts []domain.Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		return nil, err
	}
	return posts, nil
}
