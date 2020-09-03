package helper

import (
	"fmt"

	app "main/src"
	"main/src/db"
)

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
