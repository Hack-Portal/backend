package transaction

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func initRoom(t *testing.T) (string, string) {
	_, owner, room := randomRoom(t)
	_, err := store.InitChatRoom(context.Background(), room.RoomID)
	require.NoError(t, err)
	return owner.AccountID, room.RoomID
}

func TestInitRoom(t *testing.T) {
	initRoom(t)
}

// func TestWriteFireBase(t *testing.T) {
// 	ownerID, roomID := initRoom(t)

// 	docs, err := store.ReadDocsByRoomID(context.Background(), roomID)
// 	require.NoError(t, err)

// 	arg := domain.WriteFireStoreParam{
// 		RoomID:  roomID,
// 		Index:   len(docs) + 1,
// 		UID:     ownerID,
// 		Message: "testing",
// 	}

// 	result, err := store.Crea(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, result)

// 	newDocs, err := store.ReadDocsByRoomID(context.Background(), roomID)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, newDocs)
// 	require.Len(t, newDocs, 1)
// }

func TestUploadImage(t *testing.T) {
	image, err := ioutil.ReadFile("../../../../../color.png")
	require.NoError(t, err)
	require.NotEmpty(t, image)

	filename, path, err := store.UploadImage(context.Background(), image)
	require.NoError(t, err)
	require.NotEmpty(t, path)

	require.NoError(t, store.DeleteImage(context.Background(), filename+".jpg"))
}
