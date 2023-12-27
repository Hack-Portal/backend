package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/Hack-Portal/backend/cmd/migrations"
	"github.com/Hack-Portal/backend/src/driver/aws"
	"github.com/Hack-Portal/backend/src/frameworks/db/gorm"
	"github.com/Hack-Portal/backend/src/frameworks/echo"
	"github.com/Hack-Portal/backend/src/server"
	"github.com/murasame29/db-conn/sqldb/postgres"
)

func init() {
	config.LoadEnv()
	// config.LoadEnv("")
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
	logger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		},
	))

	postgresConn := postgres.NewConnection(
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.DBName,
		config.Config.Database.SSLMode,
		config.Config.Database.ConnectAttempts,
		config.Config.Database.ConnectWaitTime,
		config.Config.Database.ConnectBlocks,
	)

	sqlDB, err := postgresConn.Connection()
	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect database: %v", err))
		return
	}

	// open db connection
	gorm := gorm.New()
	dbconn, err := gorm.Connection(sqlDB)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect database: %v", err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer gorm.Close(ctx)

	// migrate
	m, err := migrations.NewPostgresMigrate(sqlDB, "file://cmd/migrations", nil)
	if err != nil {
		logger.Error(fmt.Sprintf("failed create migrate instance: %v", err))
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

	// start server
	handler := echo.NewEchoServer(
		dbconn,
		client,
		logger,
	)

	server.New().Run(handler)
}
