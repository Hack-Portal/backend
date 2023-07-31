package db

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
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
	result, err := store.client.Collection(FireStoreChatRoomCollectionName).Doc(arg.RoomID).Update(ctx, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 初期化する
func (store *SQLStore) InitChatRoom(ctx context.Context, roomID string) (*firestore.WriteResult, error) {
	result, err := store.client.Collection(FireStoreChatRoomCollectionName).Doc(roomID).Set(ctx, map[string]interface{}{}, firestore.MergeAll)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// ルームIDからドキュメントを取得する
func (store *SQLStore) ReadDocsByRoomID(ctx context.Context, RoomID string) (map[string]ChatRoomsWrite, error) {
	dsnap, err := store.client.Collection(FireStoreChatRoomCollectionName).Doc(RoomID).Get(context.Background())
	if err != nil {
		return nil, err
	}
	var data map[string]ChatRoomsWrite
	dsnap.DataTo(&data)
	return data, nil
}
