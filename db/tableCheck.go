package db

import (
	"database/sql"
)

func CheckIfTableExists(db *sql.DB, tableName string) (bool, error) {
	query := "SELECT name FROM sqlite_master WHERE type='table' AND name=?;"

	var name string
	err := db.QueryRow(query, tableName).Scan(&name)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // Table does not exist
		}
		return false, err
	}

	return true, nil
}
