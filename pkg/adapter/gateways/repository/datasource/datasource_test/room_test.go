package repository_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	"github.com/stretchr/testify/require"
)

func createRoomsTest(t *testing.T) repository.Room {
	hackathon := createHackathonTest(t)

	arg := repository.CreateRoomsParams{
		RoomID:      uuid.New(),
		HackathonID: hackathon.HackathonID,
		Title:       util.RandomString(8),
		Description: util.RandomString(100),
		MemberLimit: 5,
		IsDelete:    false,
	}

	room, err := testQueries.CreateRooms(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, room)

	require.Equal(t, arg.RoomID, room.RoomID)
	require.Equal(t, arg.HackathonID, room.HackathonID)
	require.Equal(t, arg.Title, room.Title)
	require.Equal(t, arg.Description, room.Description)
	require.Equal(t, arg.MemberLimit, room.MemberLimit)
	require.Equal(t, arg.IsDelete, room.IsDelete)
	require.NotZero(t, room.CreateAt)
	return room
}

func TestCreateRoom(t *testing.T) {
	createRoomsTest(t)
}

func TestGetRoom(t *testing.T) {
	room1 := createRoomsTest(t)

	room2, err := testQueries.GetRoomsByID(context.Background(), room1.RoomID)
	require.NoError(t, err)
	require.NotEmpty(t, room2)

	require.Equal(t, room1.RoomID, room2.RoomID)
	require.Equal(t, room1.HackathonID, room2.HackathonID)
	require.Equal(t, room1.Title, room2.Title)
	require.Equal(t, room1.Description, room2.Description)
	require.Equal(t, room1.MemberLimit, room2.MemberLimit)
	require.Equal(t, room1.IsDelete, room2.IsDelete)
}

func TestListRoom(t *testing.T) {
	n := 5
	for i := 0; i < n; i++ {
		createRoomsTest(t)
	}
	arg := repository.ListRoomsParams{
		Limit:  int32(n),
		Offset: 0,
	}
	rooms, err := testQueries.ListRooms(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, rooms)
	require.Len(t, rooms, n)
	for _, room := range rooms {
		require.NotEmpty(t, room)
	}
}
