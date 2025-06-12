package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	return db, nil
}
