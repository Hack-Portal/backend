package main

import (
	"context"
	"database/sql"
	"log"

	firebase "firebase.google.com/go"
	"github.com/hackhack-Geek-vol6/backend/api"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/util"
	_ "github.com/lib/pq"
	"google.golang.org/api/option"
)

func main() {
	config, err := util.LoadEnvConfig(".")
	if err != nil {
		log.Fatal("cannnot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	serviceAccount := option.WithCredentialsFile("./serviceAccount.json")
	app, err := firebase.NewApp(context.Background(), nil, serviceAccount)

	if err != nil {
		log.Fatal("cerviceAccount Load error :", err)
	}

	store := db.NewStore(conn, app)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal(err)
	}
	if err := server.Start(config.ServerPort); err != nil {
		log.Fatal("cannnot start server :", err)
	}
}
