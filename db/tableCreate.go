package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/theSC0RP/cli-todo/utils"
)

func CreateTableIfNotExists(db *sql.DB, tableName string, columns map[string]string) error {
	if !utils.IsValidTableIdentifier(tableName) {
		return fmt.Errorf("invalid table name: %s", tableName)
	}

	columnDefs := make([]string, 0, len(columns))
	for columnName, columnType := range columns {
		if !utils.IsValidTableIdentifier(columnName) {
			return errors.New("invalid column name: " + columnName)
		}
		columnDefs = append(columnDefs, columnName+" "+columnType)
	}

	queryBuilder := strings.Builder{}
	queryBuilder.WriteString("CREATE TABLE IF NOT EXISTS ")
	queryBuilder.WriteString(tableName)
	queryBuilder.WriteString(" (")
	queryBuilder.WriteString(strings.Join(columnDefs, ", "))
	queryBuilder.WriteString(");")
	query := queryBuilder.String()

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
