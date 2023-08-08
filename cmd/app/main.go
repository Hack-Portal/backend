package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	v1 "github.com/hackhack-Geek-vol6/backend/pkg/gateways/infrastructure/httpserver/route/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/gateways/repository/transaction"
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
	env := bootstrap.LoadEnvConfig(".")
	db, err := sql.Open(env.DBDriver, env.DBSource)
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

	store := transaction.NewStore(db, app)

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()
	v1.Setup(&env, timeout, store, gin)

	gin.Run(env.ServerPort)
}
