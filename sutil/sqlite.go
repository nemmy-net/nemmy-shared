package sutil

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// How every DB must be opened.
// This ensures that tests will experience the same DB errors.
func OpenDb(filename string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", filename+"?mode=rwc&_journal_mode=WAL&_synchronous=NORMAL")
	if err != nil {
		return db, err
	}

	// Required for SQLite3 on Linux to prevent "database is locked" errors
	db.SetMaxIdleConns(2)
	return db, err
}
