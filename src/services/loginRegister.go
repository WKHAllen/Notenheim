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

	sql = `
		INSERT INTO AppUser
			(id, email, password, verified, joinTimestamp)
		VALUES
			(?, ?, ?, FALSE, ?);`
	hashed, err := app.HashPassword(password)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return fmt.Errorf("An unexpected error occurred")
	}

	err = dbm.Execute(sql, app.Base64ID(4), email, hashed, app.GetTime())
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return fmt.Errorf("An unexpected error occurred")
	}

	return nil
}
