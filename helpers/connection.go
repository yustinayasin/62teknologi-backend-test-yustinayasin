package helpers

import (
	"database/sql"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() (*Database, error) {
	//connection
	connStr := "user=postgres dbname=62teknologi password=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}
