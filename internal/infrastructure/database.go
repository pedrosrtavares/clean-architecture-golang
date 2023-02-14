package infrastructure

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDatabaseConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/clean_architecture_golang?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	return db
}
