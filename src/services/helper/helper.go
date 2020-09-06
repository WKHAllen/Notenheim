package helper

import (
	"fmt"

	app "main/src"
	"main/src/db"
)

// UnexpectedError returns a generic error message back and prints a more helpful message
func UnexpectedError(dbm *db.Manager, sql string, params ...interface{}) error {
	err := dbm.Execute(sql, params...)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return fmt.Errorf("An unexpected error occurred")
	}
	return nil
}

// UniqueBase64ID generates a new random ID in base 64 unique to the column it will go in
func UniqueBase64ID(length int, dbm *db.Manager, table string, column string) string {
	id := app.Base64ID(length)
	var dupID string

	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?;", column, table, column)
	err := dbm.QueryRow(sql, id).Scan(&dupID)
	if err == nil {
		return UniqueBase64ID(length, dbm, table, column)
	}

	return id
}

// StructureRow organizes a row into a logical key-value structure
func StructureRow(row []interface{}, keys ...string) map[string]interface{} {
	structuredRow := make(map[string]interface{})

	for index, key := range keys {
		structuredRow[key] = row[index]
	}

	return structuredRow
}
