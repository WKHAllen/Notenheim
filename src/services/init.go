package services

import (
	"main/src/db"
)

var dbm *db.Manager

// SetDBManager sets the package's global database manager
func SetDBManager(dbManager *db.Manager) {
	dbm = dbManager
}
