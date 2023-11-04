package firebase

import (
	"context"

	fb "firebase.google.com/go"
	"google.golang.org/api/option"
)

func FbApp(path string) (*fb.App, error) {
	firebaseconfig := &fb.Config{
		StorageBucket: "hack-portal-2.appspot.com",
	}

	serviceAccount := option.WithCredentialsFile(path)
	app, err := fb.NewApp(context.Background(), firebaseconfig, serviceAccount)
	if err != nil {
		return nil, err
	}

	return app, nil
}
