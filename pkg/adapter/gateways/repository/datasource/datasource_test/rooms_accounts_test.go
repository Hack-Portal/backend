package repository_test

import (
	"context"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/stretchr/testify/require"
)

func createRoomsAccountsTest(t *testing.T, room repository.Room, account repository.Account) repository.RoomsAccount {

	arg := repository.CreateRoomsAccountsParams{
		UserID:  account.UserID,
		RoomID:  room.RoomID,
		IsOwner: false,
	}
	roomsAccounts, err := testQueries.CreateRoomsAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, roomsAccounts)

	require.Equal(t, arg.RoomID, roomsAccounts.RoomID)
	require.Equal(t, arg.UserID, roomsAccounts.UserID)
	return roomsAccounts
}

func TestCreateRoomsAccounts(t *testing.T) {
	room := createRoomsTest(t)
	account := createAccountTest(t)
	createRoomsAccountsTest(t, room, account)
}

func TestGetRoomsAccounts(t *testing.T) {
	room := createRoomsTest(t)
	account := createAccountTest(t)
	n := 5

	for i := 0; i < n; i++ {
		createRoomsAccountsTest(t, room, account)
	}

	tags, err := testQueries.GetRoomsAccountsByID(context.Background(), room.RoomID)
	require.NoError(t, err)
	require.NotEmpty(t, tags)
	require.Len(t, tags, n)
}

// ToDo:時間のあるとき実装する
// func TestRemoveAccountInRoom(t *testing.T) {
// 	room := createRoomsTest(t)
// 	account := createAccountTest(t)
// 	roomAccount := createRoomsAccountsTest(t, room, account)
// 	log.Println("1 :", roomAccount)

// 	roomAccounts, err := testQueries.RemoveAccountInRoom(context.Background(), RemoveAccountInRoomParams{
// 		RoomID: room.RoomID,
// 		UserID: account.UserID,
// 	})
// 	log.Println("2 :", roomAccounts)
// 	require.NoError(t, err)
// 	require.Empty(t, roomAccounts)
// }
