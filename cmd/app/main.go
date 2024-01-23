package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/Hack-Portal/backend/cmd/migrations"
	"github.com/Hack-Portal/backend/src/driver/aws"
	"github.com/Hack-Portal/backend/src/driver/newrelic"
	"github.com/Hack-Portal/backend/src/driver/redis"
	"github.com/Hack-Portal/backend/src/router"
	"github.com/Hack-Portal/backend/src/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	config.LoadEnv()
	//config.LoadEnv(".env")
}

//	@title						Hack-Portal Backend API
//	@version					0.1.0
//	@description			Hack-Portal Backend API serice
//	@termsOfService	　https://hc-dev.seafood-dev.com

//	@contact.name			murasame29
//	@contact.url			https://twitter.com/fresh_salmon256
//	@contact.email		oogiriminister@gmail.com

//	@license.name			No-license

// @host							api-dev.hack-portal.com
// @BasePath					/v1
func main() {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.DBName,
		config.Config.Database.SSLMode,
	)
	log.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("gorm open error: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("db.DB error: ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("db ping error: ", err)
	}

	// migrate
	m, err := migrations.NewPostgresMigrate(sqlDB, "file://cmd/migrations", nil)
	if err != nil {
		log.Fatal("migrate error: ", err)
	}

	// migrate up
	log.Println(m.Up())

	client, err := aws.New(
		config.Config.Buckets.AccountID,
		config.Config.Buckets.EndPoint,
		config.Config.Buckets.AccessKeyId,
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

	redisconn := redis.New(
		fmt.Sprintf("%v:%v", config.Config.Redis.Host, config.Config.Redis.Port),
		config.Config.Redis.Password,
		&config.Config.Redis.ConnectTimeout,
		&config.Config.Redis.ConnectAttempts,
		&config.Config.Redis.ConnectWaitTime,
	)
	defer redisconn.Close()

	redisConn, err := redisconn.Connect(config.Config.Redis.DB)
	if err != nil {
		log.Fatal(err)
	}

	// start server
	handler := router.NewRouter(
		router.NewDebug(config.Config.Server.Version),
		db,
		redisConn,
		nrapp,
		client,
	)

	server.New().Run(handler)
}
