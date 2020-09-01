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
	// Person table
	panicExec(dbm, `DROP TABLE IF EXISTS person;`)
	panicExec(dbm, `
		CREATE TABLE IF NOT EXISTS person (
			id        SERIAL PRIMARY KEY NOT NULL,
			firstname VARCHAR(255)       NOT NULL,
			lastname  VARCHAR(255)       NOT NULL
		);
	`)
	people := []struct{
		first string
		last  string
	}{
		{first: "Jonas",  last: "Kahnwald"},
		{first: "Martha", last: "Nielsen"},
		{first: "Hanno",  last: "Tauber"},
	}
	for _, person := range people {
		panicExec(dbm, `INSERT INTO person (firstname, lastname) VALUES (?, ?);`, person.first, person.last)
	}

	// Message table
	panicExec(dbm, `DROP TABLE IF EXISTS quote;`)
	panicExec(dbm, `
		CREATE TABLE IF NOT EXISTS quote (
			id   SERIAL PRIMARY KEY NOT NULL,
			text VARCHAR(1024)      NOT NULL
		);
	`)
	quotes := []string{
		"What we know is a drop. What we don't know is an ocean.",
		"We're not free in what we do because we're not free in what we want.",
		"Yesterday, today and tomorrow are not consecutive, they are connected in a never-ending circle.",
		"A man lives three lives. The first one ends with the loss of naivety, the second, with the loss of innocence and the third with the loss of life itself.",
		"Only when we’ve freed ourselves of emotion can we be truly free. Only when you’re willing to sacrifice what you hold dearest.",
		"The beginning is the end, and the end is the beginning.",
		"We all face the same end. Those above have forgotten us. They do not judge us. In death I am alone, and my only judge is me.",
	}
	for _, quote := range quotes {
		panicExec(dbm, `INSERT INTO quote (text) VALUES (?);`, quote)
	}
}
