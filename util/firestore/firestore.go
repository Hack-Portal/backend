package firestore

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// クライアントをセットする
func FirebaseSetup(path string) (*firebase.App, error) {
	config := &firebase.Config{
		StorageBucket: "hackthon-geek-v6.appspot.com",
	}

	serviceAccount := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(context.Background(), config, serviceAccount)
	if err != nil {
		return nil, err
	}
	return app, nil
}
