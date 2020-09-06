package services

import (
	"fmt"

	app "main/src"
	"main/src/services/helper"
)

// NewListItem creates a new item in a list
func NewListItem(sessionID string, listID string) (string, error) {
	var title string

	// Confirm that the session ID exists
	userID, err := helper.GetUserSession(dbm, sessionID)
	if err != nil { return "", err }

	// Confirm that the list ID is valid
	sql := "SELECT title FROM List WHERE id = ? AND userID = ?;"
	err = dbm.QueryRow(sql, listID, userID).Scan(&title)
	if err != nil {
		return "", fmt.Errorf("Invalid list ID")
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

// EditListItem edits a list item's content
func EditListItem(sessionID string, listItemID string, newContent string) error {
	var listID string
	var title string

	// Confirm that the session ID exists
	userID, err := helper.GetUserSession(dbm, sessionID)
	if err != nil { return err }

	// Confirm that the list item ID is valid
	sql := "SELECT listID FROM ListItem WHERE id = ?;"
	err = dbm.QueryRow(sql, listItemID).Scan(&listID)
	if err != nil {
		return fmt.Errorf("Invalid list item ID")
	}

	// Confirm that the list ID is valid
	sql = "SELECT title FROM List WHERE id = ? AND userID = ?;"
	err = dbm.QueryRow(sql, listID, userID).Scan(&title)
	if err != nil {
		return fmt.Errorf("Invalid list item ID")
	}

	// Edit list item content
	sql = "UPDATE ListItem SET content = ? WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, newContent, listItemID)
	if err != nil { return err }

	return nil
}

// CheckListItem checks or unchecks a list item
func CheckListItem(sessionID string, listItemID string, checked bool) error {
	var listID string
	var title string

	// Confirm that the session ID exists
	userID, err := helper.GetUserSession(dbm, sessionID)
	if err != nil { return err }

	// Confirm that the list item ID is valid
	sql := "SELECT listID FROM ListItem WHERE id = ?;"
	err = dbm.QueryRow(sql, listItemID).Scan(&listID)
	if err != nil {
		return fmt.Errorf("Invalid list item ID")
	}

	// Confirm that the list ID is valid
	sql = "SELECT title FROM List WHERE id = ? AND userID = ?;"
	err = dbm.QueryRow(sql, listID, userID).Scan(&title)
	if err != nil {
		return fmt.Errorf("Invalid list item ID")
	}

	// Check/uncheck the list item
	sql = "UPDATE ListItem SET checked = ? WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, checked, listItemID)
	if err != nil { return err }

	return nil
}
