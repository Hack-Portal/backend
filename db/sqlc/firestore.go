package db

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

type WriteFireStoreParam struct {
	RoomID  string `json:"room_id"`
	Index   int    `json:"index"`
	UID     string `json:"uid"`
	Message string `json:"message"`
}

type ChatRoomsWrite struct {
	UID       string    `json:"uid"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

const (
	FireStoreChatRoomCollectionName = "chatrooms"
)

// fireStoreにデータを追加する
func (store *SQLStore) WriteFireStore(ctx context.Context, arg WriteFireStoreParam) (*firestore.WriteResult, error) {
	update := []firestore.Update{
		{
			Path: fmt.Sprint(arg.Index),
			Value: ChatRoomsWrite{
				UID:       arg.UID,
				Message:   arg.Message,
				CreatedAt: time.Now(),
			},
		},
	}
	client, err := store.client.Firestore(ctx)
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
	client, err := store.client.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	result, err := client.Collection(FireStoreChatRoomCollectionName).Doc(roomID).Set(ctx, map[string]interface{}{}, firestore.MergeAll)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// ルームIDからドキュメントを取得する
func (store *SQLStore) ReadDocsByRoomID(ctx context.Context, RoomID string) (map[string]ChatRoomsWrite, error) {
	client, err := store.client.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	dsnap, err := client.Collection(FireStoreChatRoomCollectionName).Doc(RoomID).Get(context.Background())
	if err != nil {
		return nil, err
	}
	var data map[string]ChatRoomsWrite
	dsnap.DataTo(&data)
	return data, nil
}

// firebaseCloudStorageに画像を上げる
func (store *SQLStore) UploadImage(ctx context.Context, file []byte, filename string) (uuid.UUID, error) {
	id := uuid.New()
	// パス取得
	fbstorage, err := store.client.Storage(ctx)
	if err != nil {
		return id, err
	}
	bucket, err := fbstorage.DefaultBucket()
	if err != nil {
		return id, err
	}

	object := bucket.Object(filename)
	writer := object.NewWriter(ctx)

	//Set the attribute
	writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
	defer writer.Close()

	if _, err := io.Copy(writer, bytes.NewReader(file)); err != nil {
		return id, err
	}

	if err := object.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
		return id, err
	}

	return id, nil
}
