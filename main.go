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

// @title           Hack Hack Backend API
// @version         1.0
// @description     HackPortal Backend API serice
// @termsOfService  https://seaffood.com/api

// @contact.name   murasame
// @contact.url    https://twitter.com/fresh_salmon256
// @contact.email  oogiriminister@gmail.com

// @license.name  No-license
// @license.url   No-license

// @host      https://seaffood.com
// @BasePath  /api/v1

func main() {
	config, err := util.LoadEnvConfig(".")
	if err != nil {
		log.Fatal("cannnot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	firebaseconfig := &firebase.Config{
		StorageBucket: "hackthon-geek-v6.appspot.com",
	}

	serviceAccount := option.WithCredentialsFile("./serviceAccount.json")
	app, err := firebase.NewApp(context.Background(), firebaseconfig, serviceAccount)

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
