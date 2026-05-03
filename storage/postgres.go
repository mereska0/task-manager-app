package storage

import (
	"database/sql"
	"log"
)

func NewPostgresDB() *sql.DB {
	connStr := "host=localhost port=5432 user=postgres password=1234 dbname=tmanager sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Create table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		text TEXT NOT NULL,
		isdone BOOLEAN NOT NULL DEFAULT FALSE
	)`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	return db
}
