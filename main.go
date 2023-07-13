package main

import (
	"database/sql"
	"log"

	"github.com/hackhack-Geek-vol6/backend/api"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/hackhack-Geek-vol6/backend/util/firestore"
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
	client, err := firestore.FirebaseSetup("./serviceAccount.json")
	if err != nil {
		log.Fatal("firestore error :", err)
	}

	store := db.NewStore(conn, client)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal(err)
	}
	if err := server.Start("0.0.0.0:8080"); err != nil {
		log.Fatal("cannnot start server :", err)
	}
}
