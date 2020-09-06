package services

import (
	"fmt"

	app "main/src"
	"main/src/services/helper"
)

// NewListItem creates a new item in a list
func NewListItem(sessionID string, listID string) (string, error) {
	var userID string

	// Confirm that the session ID exists
	sql := "SELECT userID FROM Session WHERE id = ?;"
	err := dbm.QueryRow(sql, sessionID).Scan(&userID)
	if err != nil {
		return "", fmt.Errorf("Invalid session")
	}

	// Add list record
	listItemID := helper.UniqueBase64ID(4, dbm, "ListItem", "id")
	now := app.GetTime()
	sql = `
		INSERT INTO ListItem
			(id, listID, content, position, checked, createTimestamp, updateTimestamp)
		VALUES
			(?, ?, ?, (
				SELECT COUNT(id) FROM ListItem WHERE listID = ?
			), ?, ?, ?);`
	err = helper.UnexpectedError(dbm, sql, listItemID, listID, "", listID, false, now, now)
	if err != nil { return "", err }

	return listItemID, nil
}
