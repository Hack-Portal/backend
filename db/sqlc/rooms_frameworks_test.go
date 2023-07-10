package db

import (
	"context"
	"testing"

	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/stretchr/testify/require"
)

func createRoomsFrameworksTest(t *testing.T, room Rooms) RoomsFrameworks {
	frameworks := listFrameworkTest(t)
	randomId := util.Random(len(frameworks) - 1)

	arg := CreateRoomsFrameworkParams{
		RoomID:      room.RoomID,
		FrameworkID: frameworks[randomId].FrameworkID,
	}

	roomsFrameworks, err := testQueries.CreateRoomsFramework(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, roomsFrameworks)

	require.Equal(t, arg.RoomID, roomsFrameworks.RoomID)
	require.Equal(t, arg.FrameworkID, roomsFrameworks.FrameworkID)

	return roomsFrameworks
}

func TestCreateRoomsFrameworks(t *testing.T) {
	room := createRoomsTest(t)
	createRoomsFrameworksTest(t, room)
}

func TestListRoomsFrameworks(t *testing.T) {
	n := 5
	room := createRoomsTest(t)
	for i := 0; i < n; i++ {
		createRoomsFrameworksTest(t, room)
	}

	listRoomsFrameworks, err := testQueries.ListRoomsFrameworks(context.Background(), room.RoomID)
	require.NoError(t, err)
	require.NotEmpty(t, listRoomsFrameworks)
	require.Len(t, listRoomsFrameworks, n)

	for _, accountsFramework := range listRoomsFrameworks {
		require.NotEmpty(t, accountsFramework)
	}
}
