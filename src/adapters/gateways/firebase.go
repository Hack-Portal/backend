package gateways

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/dai"
)

type FirebaseRepository struct {
	app *firebase.App
}

func NewFirebaseRepository(app *firebase.App) dai.FirebaseRepository {
	return &FirebaseRepository{app}
}

func (r *FirebaseRepository) connBucket(ctx context.Context) (*storage.BucketHandle, error) {
	fbstorage, err := r.app.Storage(ctx)
	if err != nil {
		return nil, err
	}
	bucket, err := fbstorage.DefaultBucket()
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

func (r *FirebaseRepository) UploadFile(hackathonID string, image []byte) (string, error) {
	// パス取得
	ctx := context.Background()

	bucket, err := r.connBucket(ctx)
	if err != nil {
		return "", err
	}

	obj := bucket.Object(hackathonID + ".jpg")
	wc := obj.NewWriter(ctx)
	wc.ContentType = "image/jpg"

	if _, err := wc.Write(image); err != nil {
		return "", fmt.Errorf("createFile:write %v: %v", hackathonID, err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("createFile:close %v: %v", hackathonID, err)
	}

	downloadURL, err := bucket.SignedURL(obj.ObjectName(), &storage.SignedURLOptions{
		Expires: time.Now().AddDate(100, 0, 0),
		Method:  "GET",
	})

	if err != nil {
		return "", fmt.Errorf("createFile:signedURL %v: %v", hackathonID, err)
	}

	return downloadURL, nil
}

func (r *FirebaseRepository) DeleteFile(hackathonID string) error {
	ctx := context.Background()
	bucket, err := r.connBucket(ctx)
	if err != nil {
		return err
	}

	if err := bucket.Object(hackathonID + ".jpg").Delete(ctx); err != nil {
		return err
	}

	return nil
}
