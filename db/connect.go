package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", ".cli-todo.db")
	return db, err
}

func CloseConnection(db *sql.DB) {
	db.Close()
}
