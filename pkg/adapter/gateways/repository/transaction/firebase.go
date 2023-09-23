package transaction

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"github.com/gofrs/uuid/v5"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/params"
	"google.golang.org/api/iterator"
)

const (
	FireStoreChatRoomCollectionName = "chatrooms"
	FireStoreSubCollectionName      = "chats"
)

type Chat struct {
	UID       string
	Message   string
	CreatedAt time.Time
}

func (store *SQLStore) CreateSubCollection(ctx context.Context, arg params.WriteFireStore) (*firestore.WriteResult, error) {
	client, err := store.App.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return client.Collection(FireStoreChatRoomCollectionName).Doc(arg.RoomID).Collection(FireStoreSubCollectionName).Doc(fmt.Sprintf("chat%05d", arg.Index)).Set(ctx, Chat{
		UID:       arg.UID,
		Message:   arg.Message,
		CreatedAt: time.Now(),
	})
}

// 初期化する
func (store *SQLStore) InitChatRoom(ctx context.Context, roomID string) (*firestore.WriteResult, error) {
	client, err := store.App.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	result, err := client.Collection(FireStoreChatRoomCollectionName).Doc(roomID).Collection(FireStoreSubCollectionName).Doc("init").Set(ctx, map[string]interface{}{})

	if err != nil {
		return nil, err
	}

	return result, nil
}

// ルームIDからドキュメントを取得する
func (store *SQLStore) ReadDocsByRoomID(ctx context.Context, roomID string) (int, error) {
	var cnt int
	client, err := store.App.Firestore(ctx)
	if err != nil {
		return 0, err
	}
	iter := client.Collection(FireStoreChatRoomCollectionName).Doc(roomID).Collection(FireStoreSubCollectionName).Documents(ctx)
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, err
		}
		cnt++
	}
	return cnt, nil
}

// firebaseCloudStorageに画像を上げる
func (store *SQLStore) UploadImage(ctx context.Context, file []byte) (string, string, error) {
	filename, err := uuid.NewGen().NewV7()
	print("test1")
	if err != nil {
		return "", "", err
	}
	// パス取得
	fbstorage, err := store.App.Storage(ctx)
	if err != nil {
		return "", "", err
	}
	bucket, err := fbstorage.DefaultBucket()
	if err != nil {
		return "", "", err
	}

	obj := bucket.Object(filename.String() + ".jpg")
	wc := obj.NewWriter(ctx)
	wc.ContentType = "image/jpg"

	if _, err := wc.Write(file); err != nil {
		return "", "", fmt.Errorf("createFile:file %v: %v", filename, err)
	}
	if err := wc.Close(); err != nil {
		return "", "", fmt.Errorf("createFile:file %v: %v", filename, err)
	}
	downloadURL, err := bucket.SignedURL(obj.ObjectName(), &storage.SignedURLOptions{
		Expires: time.Now().AddDate(100, 0, 0),
		Method:  "GET",
	})

	if err != nil {
		return "", "", fmt.Errorf("downloadURL :%v", err)
	}

	return filename.String(), downloadURL, nil
}

func (store *SQLStore) DeleteImage(ctx context.Context, file string) error {
	fbstorage, err := store.App.Storage(ctx)
	if err != nil {
		return err
	}

	bucket, err := fbstorage.DefaultBucket()
	if err != nil {
		return err
	}

	if err := bucket.Object(file).Delete(ctx); err != nil {
		return err
	}
	return nil
}
