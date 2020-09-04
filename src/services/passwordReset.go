package services

import (
	"fmt"

	app "main/src"
	"main/src/services/helper"
)

// RequestPasswordReset requests a password reset
func RequestPasswordReset(email string) error {
	var id string

	// Confirm that the email address exists
	sql := "SELECT id FROM AppUser WHERE email = ? AND verified = TRUE;"
	err := dbm.QueryRow(sql, email).Scan(&id)
	if err != nil {
		return nil // Email address not registered
	}

	// Check that no password reset has already been requested
	sql = "SELECT id FROM PasswordReset WHERE email = ?;"
	err = dbm.QueryRow(sql, email).Scan(&id)
	if err == nil {
		return nil // Password reset has already been requested
	}

	// Create password reset request
	sql = `
		INSERT INTO PasswordReset
			(id, email, createTimestamp)
		VALUES
			(?, ?, ?);`
	id = helper.UniqueBase64ID(16, dbm, "PasswordReset", "id")
	err = helper.UnexpectedError(dbm, sql, id, email, app.GetTime())
	if err != nil { return err }

	return nil
}

// ValidPasswordResetID verifies that a password reset ID is valid
func ValidPasswordResetID(resetID string) bool {
	var email string

	// Check if password reset ID exists
	sql := "SELECT email FROM PasswordReset WHERE id = ?;"
	err := dbm.QueryRow(sql, resetID).Scan(&email)
	return err == nil
}

// ResetPassword resets a user's password
func ResetPassword(resetID string, newPassword string) error {
	var email string

	// Confirm that the password reset ID exists
	sql := "SELECT email FROM PasswordReset WHERE id = ?;"
	err := dbm.QueryRow(sql, resetID).Scan(&email)
	if err != nil {
		return fmt.Errorf("Invalid password reset ID")
	}

	// Delete the password reset ID
	sql = "DELETE FROM PasswordReset WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, resetID)
	if err != nil { return err }

	// Hash the new password
	hashed, err := app.HashPassword(newPassword)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return fmt.Errorf("An unexpected error occurred")
	}

	// Change the password
	sql = "UPDATE AppUser SET password = ? WHERE email = ?;"
	err = helper.UnexpectedError(dbm, sql, hashed, email)
	if err != nil { return err }

	return nil
}
