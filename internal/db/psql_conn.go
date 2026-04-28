package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func NewDB() *sql.DB{
	connStr := "postgres://postgres:postgres@db:5432/marketplace?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}