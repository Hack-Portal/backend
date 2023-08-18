package transaction

import (
	"context"
	"database/sql"
	"log"
	"testing"

	firebase "firebase.google.com/go"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"google.golang.org/api/option"
)

var TestStore fake.fakeQuerier

func NewTestServer(t *testing.T) {
	env := bootstrap.LoadEnvConfig("../../../../../")
	db, err := sql.Open(env.DBDriver, env.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	firebaseconfig := &firebase.Config{
		StorageBucket: "hackthon-geek-v6.appspot.com",
	}

	serviceAccount := option.WithCredentialsFile("./serviceAccount.json")
	app, err := firebase.NewApp(context.Background(), firebaseconfig, serviceAccount)
	if err != nil {
		log.Fatal("cerviceAccount Load error :", err)
	}
	TestStore = NewStore(db, app)
}
