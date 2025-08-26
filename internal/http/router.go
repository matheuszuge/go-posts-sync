package httpx

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/matheuszuge/golang/internal/service"
)

func NewRouter(userSvc service.UserService, postSvc service.PostService) http.Handler {
	r := chi.NewRouter()
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/users", listUsers(userSvc))
	r.Get("/posts", ListPosts(postSvc))

	return r
}

func ListPosts(svc service.PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := svc.ListPosts()
		if err != nil {
			http.Error(w, "something broke", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(posts)
	}
}

func listUsers(svc service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := svc.ListUsers()
		if err != nil {
			http.Error(w, "something broke", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(users)
	}
}
