package util

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func NewSQLite3Connection() (*sql.DB) {
	dbDir := fmt.Sprintf("%s/currency.db", os.Getenv("DB_DIR"))
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		dbDir = "./currency.db"	
	}

	database, err := sql.Open("sqlite3", dbDir)
	if err != nil {
		log.Fatal(err)
	}

	initializeDatabase(database)

	return database
}

func initializeDatabase(database *sql.DB) {

	statement, err := database.Prepare(`
		CREATE TABLE IF NOT EXISTS currency (
			id INTEGER PRIMARY KEY, 
			name TEXT, 
			code TEXT, 
			codein TEXT, 
			bid FLOAT NOT NULL, 
			timestamp DATETIME
		)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}
}