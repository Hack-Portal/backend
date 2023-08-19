package transaction

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	firebase "firebase.google.com/go"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	_ "github.com/lib/pq"
	"google.golang.org/api/option"
)

var (
	testQueries *repository.Queries
	testDB      *sql.DB
	App         *firebase.App
)

func TestMain(m *testing.M) {
	var err error
	config := bootstrap.LoadEnvConfig("../../../../../")

	firebaseconfig := &firebase.Config{
		StorageBucket: "hackthon-geek-v6.appspot.com",
	}

	serviceAccount := option.WithCredentialsFile("./serviceAccount.json")
	app, err := firebase.NewApp(context.Background(), firebaseconfig, serviceAccount)
	if err != nil {
		log.Fatal("cerviceAccount Load error :", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	testQueries = repository.New(testDB)
	App = app
	os.Exit(m.Run())
}
