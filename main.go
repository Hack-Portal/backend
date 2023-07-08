package main

import (
	"database/sql"
	"log"

	"github.com/hackhack-Geek-vol6/backend/api"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadEnvConfig(".")
	if err != nil {
		log.Fatal("cannnot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSouse)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal(err)
	}
	if err := server.Start("0.0.0.0:5000"); err != nil {
		log.Fatal("cannnot start server :", err)
	}
}
