package transaction

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"github.com/gofrs/uuid/v5"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
)

const (
	FireStoreChatRoomCollectionName = "chatrooms"
)

// fireStoreにデータを追加する
func (store *SQLStore) WriteFireStore(ctx context.Context, arg domain.WriteFireStoreParam) (*firestore.WriteResult, error) {
	update := []firestore.Update{
		{
			Path: fmt.Sprint(arg.Index),
			Value: domain.ChatRoomsWrite{
				UID:       arg.UID,
				Message:   arg.Message,
				CreatedAt: time.Now(),
			},
		},
	}
	client, err := store.App.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	result, err := client.Collection(FireStoreChatRoomCollectionName).Doc(arg.RoomID).Update(ctx, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 初期化する
func (store *SQLStore) InitChatRoom(ctx context.Context, roomID string) (*firestore.WriteResult, error) {
	client, err := store.App.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	result, err := client.Collection(FireStoreChatRoomCollectionName).Doc(roomID).Set(context.Background(), map[string]interface{}{}, firestore.MergeAll)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// ルームIDからドキュメントを取得する
func (store *SQLStore) ReadDocsByRoomID(ctx context.Context, RoomID string) (map[string]domain.ChatRoomsWrite, error) {
	client, err := store.App.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	dsnap, err := client.Collection(FireStoreChatRoomCollectionName).Doc(RoomID).Get(context.Background())
	if err != nil {
		return nil, err
	}
	var data map[string]domain.ChatRoomsWrite
	dsnap.DataTo(&data)
	return data, nil
}

// firebaseCloudStorageに画像を上げる
func (store *SQLStore) UploadImage(ctx context.Context, file []byte) (string, error) {
	filename, err := uuid.NewGen().NewV7()
	if err != nil {
		return "", err
	}
	// パス取得
	fbstorage, err := store.App.Storage(ctx)
	log.Println("1 :", err)
	if err != nil {
		return "", err
	}
	bucket, err := fbstorage.DefaultBucket()
	log.Println("2 :", err)
	if err != nil {
		return "", err
	}

	obj := bucket.Object(filename.String() + ".jpg")
	wc := obj.NewWriter(ctx)
	wc.ContentType = "image/jpg"

	if _, err := wc.Write(file); err != nil {
		return "", fmt.Errorf("createFile:file %v: %v", filename, err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("createFile:file %v: %v", filename, err)
	}
	downloadURL, err := bucket.SignedURL(obj.ObjectName(), &storage.SignedURLOptions{
		Expires: time.Now().AddDate(100, 0, 0),
		Method:  "GET",
	})

	if err != nil {
		return "", fmt.Errorf("downloadURL :%v", err)
	}

	return downloadURL, nil
}
