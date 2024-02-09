package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

const dbFile = "clitt.sqlite"

type DB struct {
	conn *sql.DB
}

// Task table structure : id, duration, task input text from the user, category
type Task struct {
	ID       int
	Duration time.Duration
	Task     string
	Category string
}

// InitDB creates the database and the table if they do not exist.
func InitDB() error {
	if fileExists(dbFile) {
		log.Println("Database already exists")
		return nil
	}

	// Open the database, creates the file if it doesn't exist.
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer func() {
		_ = db.Close()
	}()

	if err := createTable(db); err != nil {
		return err
	}

	return nil
}

// fileExists checks the existence of a file.
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// createTable creates the task table if it does not exist.
func createTable(db *sql.DB) error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS task (
        id INTEGER PRIMARY KEY, 
        duration TIMESTAMP, 
        task TEXT, 
        category TEXT
    );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}
	return nil
}

// NewDB creates a new database connection
func NewDB() *DB {
	db, err := sql.Open("sqlite", "./clitt.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	return &DB{conn: db}
}

// Close closes the database connection
func (d *DB) Close() {
	d.conn.Close()
}
