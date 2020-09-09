package services

import (
	"fmt"

	app "main/src"
	"main/src/services/helper"
)

// NewList creates a new list
func NewList(sessionID string, title string) (string, error) {
	// Confirm that the session ID exists
	userID, err := helper.GetUserSession(dbm, sessionID)
	if err != nil { return "", err }

	// Add list record
	listID := helper.UniqueBase64ID(4, dbm, "List", "id")
	now := app.GetTime()
	sql := `
		INSERT INTO List
			(id, userID, title, createTimestamp, updateTimestamp)
		VALUES
			(?, ?, ?, ?, ?);`
	err = helper.UnexpectedError(dbm, sql, listID, userID, title, now, now)
	if err != nil { return "", err }

	return listID, nil
}

// GetLists gets all lists created by a user
func GetLists(sessionID string) ([]app.Object, error) {
	var emptyObj []app.Object

	// Confirm that the session ID exists
	userID, err := helper.GetUserSession(dbm, sessionID)
	if err != nil { return emptyObj, err }

	// Get lists
	sql := `
		SELECT id, title, updateTimestamp
		FROM List
		WHERE userID = ?
		ORDER BY updateTimestamp DESC;`
	rows, err := dbm.QueryRows(sql, userID)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return emptyObj, fmt.Errorf("An unexpected error occurred")
	}

	// Parse list information
	lists := make([]app.Object, 0)
	for rows.Next() {
		row, err := rows.Values()
		if err != nil {
			fmt.Printf("Unexpected error: %v\n", err)
			return emptyObj, fmt.Errorf("An unexpected error occurred")
		}

		structuredRow := helper.StructureRow(row, "listID", "title", "updateTimestamp")
		lists = append(lists, structuredRow)
	}

	return lists, nil
}

// ListInfo gets the title and items within a list
func ListInfo(sessionID string, listID string) (string, []app.Object, error) {
	var emptyObj []app.Object
	var title string

	// Confirm that the session ID exists
	userID, err := helper.GetUserSession(dbm, sessionID)
	if err != nil { return "", emptyObj, err }

	// Get the list title
	sql := "SELECT title FROM List WHERE id = ? AND userID = ?;"
	err = dbm.QueryRow(sql, listID, userID).Scan(&title)
	if err != nil {
		return "", emptyObj, fmt.Errorf("Invalid list ID")
	}

	// Get list items
	sql = `
		SELECT id, content, position, checked
		FROM ListItem
		WHERE listID = ?
		ORDER BY position ASC;`
	rows, err := dbm.QueryRows(sql, listID)

	// Parse list item information
	info := make([]app.Object, 0)
	for rows.Next() {
		row, err := rows.Values()
		if err != nil {
			fmt.Printf("Unexpected error: %v\n", err)
			return "", emptyObj, fmt.Errorf("An unexpected error occurred")
		}

		structuredRow := helper.StructureRow(row, "listItemID", "content", "position", "checked")
		info = append(info, structuredRow)
	}

	return title, info, nil
}

// RenameList renames a list
func RenameList(sessionID string, listID string, newName string) error {
	var title string

	// Confirm that the session ID exists
	userID, err := helper.GetUserSession(dbm, sessionID)
	if err != nil { return err }

	// Confirm that the list ID is valid
	sql := "SELECT title FROM List WHERE id = ? AND userID = ?;"
	err = dbm.QueryRow(sql, listID, userID).Scan(&title)
	if err != nil {
		return fmt.Errorf("Invalid list ID")
	}

	// Set the list title
	sql = `
		UPDATE List
		SET title = ?, updateTimestamp = ?
		WHERE id = ? AND userID = ?;`
	err = helper.UnexpectedError(dbm, sql, newName, app.GetTime(), listID, userID)
	if err != nil { return err }

	return nil
}

// DeleteList deletes a list
func DeleteList(sessionID string, listID string) error {
	var title string

	// Confirm that the session ID exists
	userID, err := helper.GetUserSession(dbm, sessionID)
	if err != nil { return err }

	// Confirm that the list ID is valid
	sql := "SELECT title FROM List WHERE id = ? AND userID = ?;"
	err = dbm.QueryRow(sql, listID, userID).Scan(&title)
	if err != nil {
		return fmt.Errorf("Invalid list ID")
	}

	// Delete the list
	sql = "DELETE FROM List WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, listID)
	if err != nil { return err }

	// Delete the list's items
	sql = "DELETE FROM ListItem WHERE listID = ?;"
	err = helper.UnexpectedError(dbm, sql, listID)
	if err != nil { return err }

	return nil
}
