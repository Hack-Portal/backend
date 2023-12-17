package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/hackhack-Geek-vol6/backend/cmd/config"
	"github.com/hackhack-Geek-vol6/backend/cmd/migrations"
	"github.com/hackhack-Geek-vol6/backend/src/frameworks/db/gorm"
	"github.com/hackhack-Geek-vol6/backend/src/frameworks/echo"
	"github.com/hackhack-Geek-vol6/backend/src/server"
)

//	@title						Hack Hack Backend API
//	@version					0.1.0
//	@description			Hack Hack Backend API serice
//	@termsOfService	　https://api.seafood-dev.com

//	@contact.name			murasame29
//	@contact.url			https://twitter.com/fresh_salmon256
//	@contact.email		oogiriminister@gmail.com

//	@license.name			No-license

// @host							api.seafood-dev.com
// @BasePath					/v1
func main() {
	logger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		},
	))

	db := gorm.NewGormConnection(
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s TimeZone=%s",
			config.Config.Database.Host,
			config.Config.Database.Port,
			config.Config.Database.User,
			config.Config.Database.DBName,
			config.Config.Database.Password,
			config.Config.Database.SSLMode,
			config.Config.Database.TimeZone,
		),
		config.Config.Database.ConnectAttempts,
		config.Config.Database.ConnectWaitTime,
		config.Config.Database.ConnectBlocks,
		logger,
	)

	// open db connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbconn, err := db.Connection()
	defer db.Close(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect database: %v", err))
		return
	}

	// get sql.DB to use in migrations
	sqldb, err := dbconn.DB()
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get sql.DB: %v", err))
		return
	}

	// migrate
	m, err := migrations.NewPostgresMigrate(sqldb, "file://cmd/migrations", nil)
	if err != nil {
		logger.Error(fmt.Sprintf("failed create migrate instance: %v", err))
		return
	}
	// migrate up
	m.Up()

	// start server
	handler := echo.NewEchoServer(
		dbconn,
		logger,
	)
	server.New().Run(handler)
}
