package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/matheuszuge/golang/internal/db"
	httpx "github.com/matheuszuge/golang/internal/http"
	"github.com/matheuszuge/golang/internal/repository"
	"github.com/matheuszuge/golang/internal/service"
)

func main() {
	database := db.Connect()
	defer database.Close()

	log.Println("API rodando... (conectada ao banco)")

	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Println("No .env file found or error loading .env file:", err)
	}

	userRepo := repository.NewInMemoryUserRepo()
	userSvc := service.NewUserService(userRepo)

	baseURL := os.Getenv("POSTS_API_URL")

	postRepo := repository.NewAPIPostRepo(baseURL, nil)
	postSvc := service.NewPostService(postRepo)

	r := httpx.NewRouter(userSvc, postSvc)

	log.Println("API rodando em: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
