package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	log.Println("Connected to DB!")
	return db, nil
}
