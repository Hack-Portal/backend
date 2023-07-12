package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// クライアントをセットする
func FirebaseSetup(path string) (*firestore.Client, error) {
	serviceAccount := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(context.Background(), nil, serviceAccount)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		return nil, err
	}
	return client, nil
}
