package services

import (
	"fmt"
	"io/ioutil"
	"time"

	app "main/src"
	"main/src/services/helper"
)

// CreateVerification creates a verification record
func CreateVerification(email string) (string, error) {
	var verifyID string

	// Confirm that the email has been registered and the user is not yet verified
	sql := "SELECT id FROM AppUser WHERE email = ? AND verified = FALSE;"
	err := dbm.QueryRow(sql, email).Scan(&verifyID)
	if err != nil {
		return "", nil // Email address not registered
	}

	// Check that no verification has already been requested
	sql = "SELECT id FROM Verify WHERE email = ?;"
	err = dbm.QueryRow(sql, email).Scan(&verifyID)
	if err == nil {
		return "", nil // Verification has already been requested
	}

	// Generate the new verification ID
	verifyID = helper.UniqueBase64ID(16, dbm, "Verify", "id")

	// Send verification email
	content, err := ioutil.ReadFile("src/emails/verify.html")
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return "", fmt.Errorf("An unexpected error occurred")
	}
	emailBody := fmt.Sprintf(string(content), verifyID)
	err = app.SendEmail(email, "Notenheim - Account Verification", emailBody)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return "", fmt.Errorf("An unexpected error occurred")
	}

	// Create verification request
	sql = `
		INSERT INTO Verify
			(id, email, createTimestamp)
		VALUES
			(?, ?, ?);`
	err = helper.UnexpectedError(dbm, sql, verifyID, email, app.GetTime())
	if err != nil { return "", err }

	return verifyID, nil
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

// PruneVerification removes a verification record
func PruneVerification(verifyID string) {
	go func() {
		var email string
		var createTimestamp int64

		// Get email and verification creation timestamp
		sql := "SELECT email, createTimestamp FROM Verify WHERE id = ?;"
		err := dbm.QueryRow(sql, verifyID).Scan(&email, &createTimestamp)
		if err != nil {
			fmt.Printf("Unexpected error: %v\n", err)
			return
		}

		// Calculate time remaining
		createdTime := time.Unix(createTimestamp, 0)
		endTime := createdTime.Add(time.Hour)
		now := time.Now()
		
		var sleepTime time.Duration
		if now.After(endTime) {
			sleepTime = 0
		} else {
			sleepTime = endTime.Sub(now)
		}

		// Wait
		time.Sleep(sleepTime)

		// Remove the verification record
		sql = "DELETE FROM Verify WHERE id = ?;"
		err = dbm.Execute(sql, verifyID)
		if err != nil {
			fmt.Printf("Unexpected error: %v\n", err)
			return
		}

		// Remove the user who created the verification request
		sql = "DELETE FROM AppUser WHERE email = ? AND verified = FALSE;"
		err = dbm.Execute(sql, email)
		if err != nil {
			fmt.Printf("Unexpected error: %v\n", err)
			return
		}
	}()
}
