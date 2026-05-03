package service

import (
	"database/sql"
	"testing"
)

func setupTestDB(t *testing.T) *sql.DB {
	connstr := "host=localhost port=5432 user=postgres password=1234 dbname=go_tasks_test sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec("DELETE FROM tasks")
	if err != nil {
		t.Fatal(err)
	}

	return db
}
