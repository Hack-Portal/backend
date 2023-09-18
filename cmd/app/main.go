package main

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/infrastructure/httpserver/route/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	_ "github.com/lib/pq"
	"google.golang.org/api/option"
)

//	@title			Hack Hack Backend API
//	@version		1.0
//	@description	HackPortal Backend API serice
//	@termsOfService	https://seaffood.com/api

//	@contact.name	murasame
//	@contact.url	https://twitter.com/fresh_salmon256
//	@contact.email	oogiriminister@gmail.com

//	@license.name	No-license
//	@license.url	No-license

//	@host		https://seaffood.com
//	@BasePath	/api/v1

func main() {
	env := bootstrap.LoadEnvConfig(".")
	db, err := sql.Open(env.DBDriver, env.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("cannot ping to db", err)
	}

	firebaseconfig := &firebase.Config{
		StorageBucket: "hack-portal-2.appspot.com",
	}

	serviceAccount := option.WithCredentialsFile("./serviceAccount.json")
	app, err := firebase.NewApp(context.Background(), firebaseconfig, serviceAccount)
	if err != nil {
		log.Fatal("cerviceAccount Load error :", err)
	}

	store := transaction.NewStore(db, app)
	times, err := strconv.Atoi(env.ContextTimeout)
	if err != nil {
		log.Fatal("invalid timeout :", err, env.ContextTimeout)
	}

	timeout := time.Duration(times) * time.Second

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	gin := gin.Default()

	route.Setup(&env, timeout, store, gin)

	gin.Run(env.ServerPort)
}
