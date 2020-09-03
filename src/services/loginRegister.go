package services

import (
	"fmt"

	app "main/src"
)

// Register registers a user
func Register(email string, password string) error {
	var id string

	sql := "SELECT id FROM AppUser WHERE email = ?;"
	err := dbm.QueryRow(sql, email).Scan(&id)
	if err == nil {
		return fmt.Errorf("Email address has already been registered")
	}

	hashed, err := app.HashPassword(password)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return fmt.Errorf("An unexpected error occurred")
	}

	sql = `
		INSERT INTO AppUser
			(id, email, password, verified, joinTimestamp)
		VALUES
			(?, ?, ?, FALSE, ?);`
	err = dbm.Execute(sql, app.Base64ID(4), email, hashed, app.GetTime())
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

	sql := "SELECT id, password FROM AppUser WHERE email = ? AND verified = TRUE;"
	err := dbm.QueryRow(sql, email).Scan(&userID, &hashed)
	if err != nil {
		return "", fmt.Errorf("Invalid login")
	}

	if !app.ComparePassword(password, hashed) {
		return "", fmt.Errorf("Invalid login")
	}

	sessionID := app.Base64ID(8)
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

	return sessionID, nil
}
