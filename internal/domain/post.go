// Package domain contains business entities and domain logic.
package domain

type Post struct {
	UserID int32  `json:"userId"`
	ID     int32  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}


