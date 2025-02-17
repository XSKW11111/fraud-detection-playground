package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetDB() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://myuser:123@localhost:5432/transactions_db?sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
