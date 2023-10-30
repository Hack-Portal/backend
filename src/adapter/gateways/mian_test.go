package gateways

import (
	"context"
	"log"
	"temp/cmd/config"
	"temp/pkg/logger"
	"temp/src/driver/db"
	"testing"

	firebase "firebase.google.com/go"
	"gorm.io/gorm"
)

var (
	l      logger.Logger
	dbconn *gorm.DB
	app    *firebase.App
)

func TestMain(m *testing.M) {
	l = logger.NewLogger(logger.DEBUG)
	config.LoadEnv(l)
	conn := db.NewConnection(l)
	defer conn.Close(context.Background())

	store, err := conn.Connection()
	if err != nil {
		log.Fatal(err)
	}
	dbconn = store

	m.Run()
}
