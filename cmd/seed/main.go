package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func main() {
	// Find database file path in db folder
	dbDir := filepath.Join(".", "db")
	dbPath := filepath.Join(dbDir, "data.sqlite3")

	// Make sure the db directory exists
	err := os.MkdirAll(dbDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Delete old database to start fresh
	os.Remove(dbPath)

	// Connect to SQLite database
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Read and run the table schema
	schemaFile, err := os.ReadFile(filepath.Join(dbDir, "schema.sql"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(string(schemaFile))
	if err != nil {
		log.Fatal(err)
	}

	// Read and insert the seed data
	seedFile, err := os.ReadFile(filepath.Join(dbDir, "seed_data.sql"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(string(seedFile))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database seeded successfully!")
}
