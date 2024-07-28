package db

import (
	"database/sql"
)

func init() {
	var db *sql.DB
	config := config.getConfig()

	db, err := sql.Open("postgres")
}
