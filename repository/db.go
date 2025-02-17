package repository

import (
	"database/sql"
)

func GetDB() *sql.DB {
	db, err := sql.Open("postgres", "POSTGRES_URL")
	if err != nil {
		panic(err)
	}
	return db
}
