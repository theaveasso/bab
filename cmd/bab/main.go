package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"theaveasso.bab/internal/api"
	"theaveasso.bab/internal/db"
	"theaveasso.bab/internal/utility"
)

func main() {
    config, err := utility.LoadConfig(".")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }
	var conn *sql.DB
	var queries db.Store

	conn, err = sql.Open(config.DBDriver, config.DSN)
	if err != nil {
		log.Fatal("error connect to db", err)
	}

	queries = db.NewStore(conn)

	server := api.NewServer(queries)

	err = server.Start(config.Address); if err != nil {
        log.Fatal("cannot start server", err)
    }
}
