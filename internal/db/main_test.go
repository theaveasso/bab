package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"theaveasso.bab/internal/utility"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
    config, err := utility.LoadConfig("../..")
	testDB, err = sql.Open(config.DBDriver, config.DSN)
	if err != nil {
		log.Fatal("error connect to db", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
