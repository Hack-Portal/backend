package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/Hack-Portal/backend/cmd/migrations"
	"github.com/Hack-Portal/backend/src/driver/aws"
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
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.DBName,
		config.Config.Database.Port,
		config.Config.Database.SSLMode,
		config.Config.Database.TimeZone,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, _ := db.DB()

	// migrate
	m, err := migrations.NewPostgresMigrate(sqlDB, "file://cmd/migrations", nil)
	if err != nil {
		return
	}
	// migrate up
	m.Up()

	client, err := aws.New(
		config.Config.Buckets.AccountID,
		config.Config.Buckets.EndPoint,
		config.Config.Buckets.AccessKeyId,
		config.Config.Buckets.AccessKeySecret,
	).Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// app, err := newrelic.Setup()
	// if err != nil {
	// 	logger.Error(fmt.Sprintf("failed to setup newrelic: %v", err))
	// 	return
	// }

	// redisconn := redis.New(
	// 	fmt.Sprintf("%v:%v", config.Config.Redis.Host, config.Config.Redis.Port),
	// 	config.Config.Redis.Password,
	// 	&config.Config.Redis.ConnectTimeout,
	// 	&config.Config.Redis.ConnectAttempts,
	// 	&config.Config.Redis.ConnectWaitTime,
	// )
	// defer redisconn.Close()

	// redisConn, err := redisconn.Connect(config.Config.Redis.DB)
	// if err != nil {
	// 	logger.Error(fmt.Sprintf("failed to connect redis: %v", err))
	// 	return
	// }

	// start server
	handler := router.NewRouter(
		router.NewDebug(config.Config.Server.Version),
		db,
		client,
	)

	server.New().Run(handler)
}
