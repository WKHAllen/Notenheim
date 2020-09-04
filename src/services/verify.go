package services

import (
	"fmt"
	app "main/src"
	"main/src/services/helper"
)

// CreateVerification creates a verification record
func CreateVerification(email string) (string, error) {
	var id string

	// Confirm that the email has been registered and the user is not yet verified
	sql := "SELECT id FROM AppUser WHERE email = ? AND verified = FALSE;"
	err := dbm.QueryRow(sql, email).Scan(&id)
	if err != nil {
		return "", nil // Email address not registered
	}

	// Check that no verification has already been requested
	sql = "SELECT id FROM Verify WHERE email = ?;"
	err = dbm.QueryRow(sql, email).Scan(&id)
	if err == nil {
		return "", nil // Verification has already been requested
	}

	// Create verification request
	sql = `
		INSERT INTO Verify
			(id, email, createTimestamp)
		VALUES
			(?, ?, ?);`
	id = helper.UniqueBase64ID(16, dbm, "Verify", "id")
	err = helper.UnexpectedError(dbm, sql, id, email, app.GetTime())
	if err != nil { return "", err }

	return id, nil
}

// Verify attempts to mark a user as verified
func Verify(verifyID string) error {
	var email string

	// Confirm that the verification ID exists
	sql := "SELECT email FROM Verify WHERE id = ?;"
	err := dbm.QueryRow(sql, verifyID).Scan(&email)
	if err != nil {
		return fmt.Errorf("Invalid verification ID")
	}

	// Delete the verification ID
	sql = "DELETE FROM Verify WHERE id = ?;"
	err = helper.UnexpectedError(dbm, sql, verifyID)
	if err != nil { return err }

	// Update the user's verification status
	sql = "UPDATE AppUser SET verified = TRUE WHERE email = ?;"
	err = helper.UnexpectedError(dbm, sql, email)
	if err != nil { return err }

	return nil
}
