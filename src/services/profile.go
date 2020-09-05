package services

import (
	"fmt"

	app "main/src"
	"main/src/services/helper"
)

// ChangePassword changes a user's password
func ChangePassword(sessionID string, newPassword string) error {
	var userID string

	// Confirm that the session ID exists
	sql := "SELECT userID FROM Session WHERE id = ?;"
	err := dbm.QueryRow(sql, sessionID).Scan(&userID)
	if err != nil {
		return fmt.Errorf("Invalid session")
	}

	// Hash the new password
	hashed, err := app.HashPassword(newPassword)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return fmt.Errorf("An unexpected error occurred")
	}

	// Change password
	sql = "UPDATE AppUser SET password = ? WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, hashed, userID)
	if err != nil { return err }

	return nil
}
