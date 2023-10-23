package main

import (
	"context"
	"time"

	firebase "firebase.google.com/go"
	"github.com/hackhack-Geek-vol6/backend/cmd/config"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	"github.com/hackhack-Geek-vol6/backend/pkg/postgres"
	_ "github.com/lib/pq"
	"google.golang.org/api/option"
)

//	@title			Hack Hack Backend API
//	@version		1.0
//	@description	HackPortal Backend API serice
//	@termsOfService	https://api-test.seafood-dev.com

//	@contact.name	murasame
//	@contact.url	https://twitter.com/fresh_salmon256
//	@contact.email	oogiriminister@gmail.com

//	@license.name	No-license
//	@license.url	No-license

// @host		https://api-test.seafood-dev.com
// @BasePath	/v1
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.Server.ContextTimeout))
	defer cancel()

	pcon := postgres.NewConnection(l)
	defer pcon.Close(ctx)

	db, err := pcon.Connection()
	if err != nil {
		l.Panicf("database connection error :%v", err)
	}

	serviceAccount := option.WithCredentialsFile("./serviceAccount.json")
	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		StorageBucket: config.Config.Firebase.StorageBucket,
	}, serviceAccount)
	if err != nil {
		l.Panicf("firebase connection error :%v", err)
	}

	// 以下サーバー起動設定

}

var l logger.Logger

func init() {
	l = logger.NewLogger(logger.DEBUG)
	config.LoadEnv(l)
}
