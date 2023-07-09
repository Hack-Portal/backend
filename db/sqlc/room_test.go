package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/stretchr/testify/require"
)

func createaRoomsTest(t *testing.T) Rooms {
	hackathon := createHackathonTest(t)

	arg := CreateRoomParams{
		RoomID:      uuid.New(),
		HackathonID: hackathon.HackathonID,
		Title:       util.RandomString(8),
		Description: util.RandomString(100),
		MemberLimit: 5,
	}

	room, err := testQueries.CreateRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, room)

	require.Equal(t, arg.RoomID, room.RoomID)
	require.Equal(t, arg.HackathonID, room.HackathonID)
	require.Equal(t, arg.Title, room.Title)
	require.Equal(t, arg.Description, room.Description)
	require.Equal(t, arg.MemberLimit, room.MemberLimit)
	return room
}

func TestCreateRoom(t *testing.T) {
	createaRoomsTest(t)
}
