package services

import (
	"fmt"

	"main/src/db"
)

var initialized bool = false
var dbm *db.Manager

// SetDBManager sets the package's global database manager
func SetDBManager(dbManager *db.Manager) {
	dbm = dbManager

	if !initialized {
		initialized = true
		pruneTables()
	}
}

// pruneTables prunes values from relevant tables
func pruneTables() {
	tables := map[string]func(string){
		"Verify":        PruneVerification,
		"PasswordReset": PrunePasswordReset,
	}
	for table, pruneFunction := range tables {
		pruneTable(table, pruneFunction)
	}
}

// pruneTable prunes values from a table
func pruneTable(table string, pruneFunction func(string)) {
	sql := fmt.Sprintf("SELECT id FROM %s;", table)
	rows, err := dbm.QueryRows(sql)
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		return
	}

	for rows.Next() {
		var id string

		err = rows.Scan(&id)
		if err != nil {
			fmt.Printf("Unexpected error: %v\n", err)
		}

		pruneFunction(id)
	}
}
