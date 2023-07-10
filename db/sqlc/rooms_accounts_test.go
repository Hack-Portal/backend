package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRoomsAccountsTest(t *testing.T, room Rooms) RoomsAccounts {
	account := createAccountTest(t)

	arg := CreateRoomsAccountsParams{
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
	createRoomsAccountsTest(t, room)
}

func TestGetRoomsAccounts(t *testing.T) {
	room := createRoomsTest(t)
	n := 5

	for i := 0; i < n; i++ {
		createRoomsAccountsTest(t, room)
	}

	tags, err := testQueries.GetRoomsAccounts(context.Background(), room.RoomID)
	require.NoError(t, err)
	require.NotEmpty(t, tags)
	require.Len(t, tags, n)
}
