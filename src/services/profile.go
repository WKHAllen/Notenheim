package services

import (
	"fmt"

	app "main/src"
	"main/src/services/helper"
)

// GetProfileInfo gets a user's profile information
func GetProfileInfo(sessionID string) (string, int64, error) {
	var email string
	var joinTimestamp int64

	// Confirm that the session ID exists
	userID, err := helper.GetUserSession(dbm, sessionID)
	if err != nil { return "", 0, err }

	// Get profile information
	sql := `
		SELECT email, joinTimestamp
		FROM AppUser
		WHERE id = ?;`
	err = dbm.QueryRow(sql, userID).Scan(&email, &joinTimestamp)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return "", 0, fmt.Errorf("An unexpected error occurred")
	}

	return email, joinTimestamp, nil
}

// ChangePassword changes a user's password
func ChangePassword(sessionID string, oldPassword string, newPassword string) error {
	var hashed string

	// Confirm that the session ID exists
	userID, err := helper.GetUserSession(dbm, sessionID)
	if err != nil { return err }

	// Get the current password
	sql := "SELECT password FROM AppUser WHERE id = ? AND verified = TRUE;"
	err = dbm.QueryRow(sql, userID).Scan(&hashed)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return fmt.Errorf("An unexpected error occurred")
	}

	// Verify the current password
	if !app.ComparePassword(oldPassword, hashed) {
		return fmt.Errorf("Invalid password")
	}

	// Hash the new password
	hashed, err = app.HashPassword(newPassword)
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
