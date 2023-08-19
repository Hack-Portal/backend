package transaction

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	firebase "firebase.google.com/go"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	_ "github.com/lib/pq"
	"google.golang.org/api/option"
)

var (
	store *SQLStore
)

func TestMain(m *testing.M) {
	var err error
	config := bootstrap.LoadEnvConfig("../../../../../")

	firebaseconfig := &firebase.Config{
		StorageBucket: "hackthon-geek-v6.appspot.com",
	}

	serviceAccount := option.WithCredentialsFile("../../../../../serviceAccount.json")
	app, err := firebase.NewApp(context.Background(), firebaseconfig, serviceAccount)
	if err != nil {
		log.Fatal("cerviceAccount Load error :", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store = NewStore(testDB, app)
	os.Exit(m.Run())
}
