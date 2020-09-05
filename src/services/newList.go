package services

import (
	"fmt"

	app "main/src"
	"main/src/services/helper"
)

// NewList creates a new list
func NewList(sessionID string, title string) (string, error) {
	var userID string

	// Confirm that the session ID exists
	sql := "SELECT userID FROM Session WHERE id = ?;"
	err := dbm.QueryRow(sql, sessionID).Scan(&userID)
	if err != nil {
		return "", fmt.Errorf("Invalid session")
	}

	listID := helper.UniqueBase64ID(4, dbm, "Session", "id")
	now := app.GetTime()
	sql = `
		INSERT INTO List
			(id, userID, title, createTimestamp, updateTimestamp)
		VALUES
			(?, ?, ?, ?, ?);`
	err = helper.UnexpectedError(dbm, sql, listID, userID, title, now, now)
	if err != nil { return "", err }

	return listID, nil
}
