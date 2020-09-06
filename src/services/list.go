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

// GetLists gets all lists created by a user
func GetLists(sessionID string) ([](map[string]interface{}), error) {
	var emptyMap [](map[string]interface{})
	var userID string

	// Confirm that the session ID exists
	sql := "SELECT userID FROM Session WHERE id = ?;"
	err := dbm.QueryRow(sql, sessionID).Scan(&userID)
	if err != nil {
		return emptyMap, fmt.Errorf("Invalid session")
	}

	sql = `
		SELECT id, title, updateTimestamp
		FROM List
		WHERE userID = ?
		ORDER BY updateTimestamp DESC;`
	rows, err := dbm.QueryRows(sql, userID)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return emptyMap, fmt.Errorf("An unexpected error occurred")
	}

	var lists [](map[string]interface{})
	for rows.Next() {
		row, err := rows.Values()
		if err != nil {
			fmt.Printf("Unexpected error: %v\n", err)
			return emptyMap, fmt.Errorf("An unexpected error occurred")
		}

		structuredRow := helper.StructureRow(row, "id", "title", "updateTimestamp")
		lists = append(lists, structuredRow)
	}

	return lists, nil
}
