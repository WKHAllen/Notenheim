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

// MoveListItem moves a list item up or down
func MoveListItem(sessionID string, listItemID string, direction string) error {
	var listID string
	var position int
	var title string
	var maxPosition int

	// Confirm that the session ID exists
	userID, err := helper.GetUserSession(dbm, sessionID)
	if err != nil { return err }

	// Confirm that the list item ID is valid
	sql := "SELECT listID, position FROM ListItem WHERE id = ?;"
	err = dbm.QueryRow(sql, listItemID).Scan(&listID, &position)
	if err != nil {
		return fmt.Errorf("Invalid list item ID")
	}

	// Confirm that the list ID is valid
	sql = "SELECT title FROM List WHERE id = ? AND userID = ?;"
	err = dbm.QueryRow(sql, listID, userID).Scan(&title)
	if err != nil {
		return fmt.Errorf("Invalid list item ID")
	}

	// Get the greatest position
	sql = "SELECT MAX(position) FROM ListItem WHERE listID = ?;"
	err = dbm.QueryRow(sql, listID).Scan(&maxPosition)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return fmt.Errorf("An unexpected error occurred")
	}

	// Set new position value
	var newPosition int
	switch (direction) {
		case "down":
			if position == 0 {
				return fmt.Errorf("Cannot move list item down")
			}
			newPosition = position - 1
		case "up":
			if position == maxPosition {
				return fmt.Errorf("Cannot move list item up")
			}
			newPosition = position + 1
	}

	// Move list item to desired position
	sql = "UPDATE ListItem SET position = ? WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, newPosition, listItemID)
	if err != nil { return err }

	// Move new position item to original item location
	sql = "UPDATE ListItem SET position = ? WHERE listID = ? AND position = ? AND id != ?;"
	err = helper.UnexpectedError(dbm, sql, position, listID, newPosition, listItemID)
	if err != nil { return err }

	return nil
}
