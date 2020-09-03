package src

import (
	"main/src/db"
)

// panicExec executes a SQL statement and panics if an error occurs
func panicExec(dbm *db.Manager, sql string, params ...interface{}) {
	err := dbm.Execute(sql, params...)
	if err != nil {
		panic(err)
	}
}

// InitDB creates and, in some cases, populates tables in the database
func InitDB(dbm *db.Manager) {
	// AppUser table
	panicExec(dbm, `
		CREATE TABLE IF NOT EXISTS AppUser (
			id            CHAR(4)      PRIMARY KEY,
			email         VARCHAR(63)  NOT NULL,
			password      VARCHAR(255) NOT NULL,
			verified      BOOLEAN      NOT NULL,
			joinTimestamp INT          NOT NULL
		);
	`)

	// List table
	panicExec(dbm, `
		CREATE TABLE IF NOT EXISTS List (
			id              CHAR(4)      PRIMARY KEY,
			userID          CHAR(4)      NOT NULL,
			title           VARCHAR(255) NOT NULL,
			createTimestamp INT          NOT NULL,
			updateTimestamp INT          NOT NULL
		);
	`)

	// ListItem table
	panicExec(dbm, `
		CREATE TABLE IF NOT EXISTS ListItem (
			id              CHAR(4)       PRIMARY KEY,
			listID          CHAR(4)       NOT NULL,
			content         VARCHAR(1023) NOT NULL,
			position        INT           NOT NULL,
			checked         BOOLEAN       NOT NULL,
			createTimestamp INT           NOT NULL,
			updateTimestamp INT           NOT NULL
		);
	`)

	// Session table
	panicExec(dbm, `
		CREATE TABLE IF NOT EXISTS Session (
			id              CHAR(16) PRIMARY KEY,
			userID          CHAR(4)  NOT NULL,
			createTimestamp INT      NOT NULL
		);
	`)

	// Verify table
	panicExec(dbm, `
		CREATE TABLE IF NOT EXISTS Verify (
			id              CHAR(16)    PRIMARY KEY,
			email           VARCHAR(63) NOT NULL,
			createTimestamp INT         NOT NULL
		);
	`)

	// PasswordReset table
	panicExec(dbm, `
		CREATE TABLE IF NOT EXISTS PasswordReset (
			id              CHAR(16)    PRIMARY KEY,
			email           VARCHAR(63) NOT NULL,
			createTimestamp INT         NOT NULL
		);
	`)
}
