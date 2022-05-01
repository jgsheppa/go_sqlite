package cmd

import "database/sql"

type DB struct {
	SQLite *sql.DB
}
