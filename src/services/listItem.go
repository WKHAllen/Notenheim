package services

import (
	"fmt"

	"main/src"
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

	// Set the list's update timestamp
	sql = "UPDATE List SET updateTimestamp = ? WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, now, listID)
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

	// Edit list item content and update timestamp
	now := app.GetTime()
	sql = "UPDATE ListItem SET content = ?, updateTimestamp = ? WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, newContent, now, listItemID)
	if err != nil { return err }

	// Set the list's update timestamp
	sql = "UPDATE List SET updateTimestamp = ? WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, now, listID)
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

	// Check/uncheck the list item and set the update timestamp
	now := src.GetTime()
	sql = "UPDATE ListItem SET checked = ?, updateTimestamp = ? WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, checked, now, listItemID)
	if err != nil { return err }

	// Set the list's update timestamp
	sql = "UPDATE List SET updateTimestamp = ? WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, now, listID)
	if err != nil { return err }

	return nil
}

// MoveListItem moves a list item up or down
func MoveListItem(sessionID string, listItemID string, newPosition int) error {
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

	// Check that the current position and new position are not the same
	if position == newPosition {
		return nil
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

	// Check that the new position is valid
	if newPosition < 0 || newPosition > maxPosition {
		return fmt.Errorf("Invalid new position")
	}

	// Determine which items need to move and in which direction
	var startPosition int
	var stopPosition int
	var moveOperator string
	if newPosition < position {
		startPosition = newPosition
		stopPosition = position
		moveOperator = "+"
	} else {
		startPosition = position
		stopPosition = newPosition
		moveOperator = "-"
	}

	// Move list item to desired position
	sql = "UPDATE ListItem SET position = ? WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, newPosition, listItemID)
	if err != nil { return err }

	// Move new position item to original item location
	sql = fmt.Sprintf("UPDATE ListItem SET position = position %s 1 WHERE listID = ? AND position >= ? AND position <= ? AND id != ?;", moveOperator)
	err = helper.UnexpectedError(dbm, sql, listID, startPosition, stopPosition, listItemID)
	if err != nil { return err }

	// Set the list's update timestamp
	now := app.GetTime()
	sql = "UPDATE List SET updateTimestamp = ? WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, now, listID)
	if err != nil { return err }

	return nil
}

// DeleteListItem deletes a list item
func DeleteListItem(sessionID string, listItemID string) error {
	var listID string
	var position int
	var title string

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

	// Delete list item
	sql = "DELETE FROM ListItem WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, listItemID)
	if err != nil { return err }

	// Alter list item positions
	sql = "UPDATE ListItem SET position = position - 1 WHERE listID = ? AND position > ?;"
	err = helper.UnexpectedError(dbm, sql, listID, position)
	if err != nil { return err }

	// Set the list's update timestamp
	now := app.GetTime()
	sql = "UPDATE List SET updateTimestamp = ? WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, now, listID)
	if err != nil { return err }

	return nil
}
