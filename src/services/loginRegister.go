package services

import (
	"fmt"

	app "main/src"
	"main/src/services/helper"
)

// Register registers a user
func Register(email string, password string) error {
	var id string

	// Confirm that the email address has not already been used
	sql := "SELECT id FROM AppUser WHERE email = ?;"
	err := dbm.QueryRow(sql, email).Scan(&id)
	if err == nil {
		return fmt.Errorf("Email address has already been registered")
	}

	// Hash the password
	hashed, err := app.HashPassword(password)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return fmt.Errorf("An unexpected error occurred")
	}

	// Add the new user to the database
	sql = `
		INSERT INTO AppUser
			(id, email, password, verified, joinTimestamp)
		VALUES
			(?, ?, ?, FALSE, ?);`
	id = helper.UniqueBase64ID(4, dbm, "AppUser", "id")
	err = dbm.Execute(sql, id, email, hashed, app.GetTime())
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return fmt.Errorf("An unexpected error occurred")
	}

	return nil
}

// Login logs a user in
func Login(email string, password string) (string, error) {
	var userID string
	var hashed string

	// Verify the provided email exists
	sql := "SELECT id, password FROM AppUser WHERE email = ? AND verified = TRUE;"
	err := dbm.QueryRow(sql, email).Scan(&userID, &hashed)
	if err != nil {
		return "", fmt.Errorf("Invalid login")
	}

	// Verify the provided password
	if !app.ComparePassword(password, hashed) {
		return "", fmt.Errorf("Invalid login")
	}

	// Create a new session for the user
	sessionID := helper.UniqueBase64ID(8, dbm, "Session", "id")
	sql = `
		INSERT INTO Session
			(id, userID, createTimestamp)
		VALUES
			(?, ?, ?);`
	err = dbm.Execute(sql, sessionID, userID, app.GetTime())
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return "", fmt.Errorf("An unexpected error occurred")
	}

	// Delete all but the user's four most recent sessions
	sql = `
		DELETE FROM Session
		WHERE userID = ? AND id NOT IN (
			SELECT id FROM Session
			WHERE userID = ?
			ORDER BY createTimestamp DESC
			LIMIT 4
		);`
	err = dbm.Execute(sql, userID, userID)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return "", fmt.Errorf("An unexpected error occurred")
	}

	return sessionID, nil
}

// Logout logs a user out
func Logout(sessionID string) error {
	// Delete the provided session
	sql := "DELETE FROM Session WHERE id = ?;"
	err := dbm.Execute(sql, sessionID)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return fmt.Errorf("An unexpected error occurred")
	}

	return nil
}
