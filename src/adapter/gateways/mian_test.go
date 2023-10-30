package gateways

import (
	"context"
	"log"
	"testing"

	firebase "firebase.google.com/go"
	"github.com/hackhack-Geek-vol6/backend/cmd/config"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	"github.com/hackhack-Geek-vol6/backend/src/driver/db"
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
