// Package db provides functions for connecting to the database.
package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	dsn := "golang:abc@123@tcp(172.24.240.1:3306)/db_golang?parseTime=true&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erro ao abrir conexão:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao conectar:", err)
	}
	return db
}
