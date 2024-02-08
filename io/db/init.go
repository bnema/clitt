package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type DB struct {
	conn *sql.DB
}

// Task table structure : id, duration, task input text from the user, category
type Task struct {
	ID       uuid.UUID
	Duration time.Duration
	Task     string
	Category string
}

// InitDB creates the database and the table
func InitDB() {
	// check if the file exists
	if _, err := os.Stat("clitt.db"); err == nil {
		log.Println("Database already exists")
		return
	} else if os.IsNotExist(err) {
		// create file if not exists
		f, err := os.Create("clitt.db")
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}
	// open the database
	db, err := sql.Open("sqlite", "./clitt.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// if the table exists, don't create it
	_, err = db.Exec("create table if not exists task (id text not null primary key, duration timestamp, task text, category text)")
	if err != nil {
		log.Fatal(err)
	}

	// success in log
	log.Println("Database initialized")
}

// NewDB creates a new database connection
func NewDB() *DB {
	db, err := sql.Open("sqlite", "./clitt.db")
	if err != nil {
		log.Fatal(err)
	}
	return &DB{conn: db}
}

// Close closes the database connection
func (d *DB) Close() {
	d.conn.Close()
}
