package transaction

import (
	"context"
	"log"
	"testing"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var TestStore Store

func NewTestServer(t *testing.T) {

	firebaseconfig := &firebase.Config{
		StorageBucket: "hackthon-geek-v6.appspot.com",
	}

	serviceAccount := option.WithCredentialsFile("./serviceAccount.json")
	app, err := firebase.NewApp(context.Background(), firebaseconfig, serviceAccount)
	if err != nil {
		log.Fatal("cerviceAccount Load error :", err)
	}
	TestStore = NewFakeStore(app)
}
