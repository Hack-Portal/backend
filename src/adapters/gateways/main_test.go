package gateways

import (
	"context"
	"log"
	"os"
	"testing"

	firebase "firebase.google.com/go"
	"github.com/hackhack-Geek-vol6/backend/cmd/config"
	"github.com/hackhack-Geek-vol6/backend/src/drivers/postgres"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	app *firebase.App
)

func TestMain(m *testing.M) {
	conn := postgres.NewConnection()
	defer conn.Close(context.Background())

	dbconn, err := conn.Connection()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db = dbconn

	firebaseconfig := &firebase.Config{
		StorageBucket: config.Config.Firebase.StorageBucket,
	}

	serviceAccount := option.WithCredentialsFile("../../../serviceAccount.json")
	fbapp, err := firebase.NewApp(context.Background(), firebaseconfig, serviceAccount)
	if err != nil {
		log.Fatal("cerviceAccount Load error :", err)
	}
	app = fbapp

	os.Exit(m.Run())
}
