package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

const (
	dbDriver = "postgres"
	dsn      = "postgres://postgres:secret@localhost:5432/bab?sslmode=disable"
)

func TestMain(m *testing.M) {
    var err error
	testDB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("error connect to db", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
