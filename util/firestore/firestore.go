package firestore

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// クライアントをセットする
func FirebaseSetup(path string) (*firebase.App, error) {
	serviceAccount := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(context.Background(), nil, serviceAccount)
	if err != nil {
		return nil, err
	}
	return app, nil
}
