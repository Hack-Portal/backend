package main

import (
	"context"
	"flag"
	"log"
	"strings"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/Hack-Portal/backend/src/driver/aws"
	"github.com/Hack-Portal/backend/src/driver/db"
	"github.com/Hack-Portal/backend/src/driver/newrelic"
	"github.com/Hack-Portal/backend/src/router"
	"github.com/Hack-Portal/backend/src/server"
	"github.com/Hack-Portal/backend/src/utils/otel"
)

type envFlag []string

func (e *envFlag) String() string {
	return strings.Join(*e, ",")
}

func (e *envFlag) Set(v string) error {
	*e = append(*e, v)
	return nil
}

func init() {
	// Usage: eg. go run main.go -e .env -e hoge.env -e fuga.env ...
	var envFile envFlag
	flag.Var(&envFile, "e", "path to .env file \n eg. -e .env -e another.env . ")
	flag.Parse()

	if err := config.LoadEnv(envFile...); err != nil {
		log.Fatal("Error loading .env file")
	}
}

//	@title						Hack-Portal Backend API
//	@version					0.1.0
//	@description			Hack-Portal Backend API serice
//	@termsOfService		https://hc-dev.seafood-dev.com

//	@contact.name			murasame29
//	@contact.url			https://twitter.com/fresh_salmon256
//	@contact.email		oogiriminister@gmail.com

//	@license.name			No-license

// @host							api-dev.hack-portal.com
// @BasePath					/v1
func main() {
	shutdown := otel.InitProvider(context.Background())
	defer shutdown()

	sqlDB := db.ConnectDBWithOtelSQL()
	defer sqlDB.Close()

	gorm := db.NewGORM(sqlDB)

	client, err := aws.New(
		config.Config.Buckets.AccountID,
		config.Config.Buckets.EndPoint,
		config.Config.Buckets.AccessKeyID,
		config.Config.Buckets.AccessKeySecret,
	).Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	nrapp, err := newrelic.Setup()
	if err != nil {
		log.Fatal(err)
		return
	}

	// start server
	handler := router.NewRouter(
		router.NewDebug(config.Config.Server.Version),
		gorm,
		nrapp,
		client,
	)

	server.New().Run(handler)
}
