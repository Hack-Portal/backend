package db

import (
	"context"
	"testing"

	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/stretchr/testify/require"
)

func createRoomsTechTagTest(t *testing.T, room Rooms) RoomsTechTags {
	tags := listTechTagTest(t)
	techs := util.Random(len(tags) - 1)

	arg := CreateRoomsTechTagParams{
		RoomID:    room.RoomID,
		TechTagID: tags[techs].TechTagID,
	}

	techTag, err := testQueries.CreateRoomsTechTag(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, techTag)

	require.Equal(t, arg.RoomID, techTag.RoomID)
	require.Equal(t, arg.TechTagID, techTag.TechTagID)
	return techTag
}

func TestCreateRoomTechTag(t *testing.T) {
	room := createRoomsTest(t)
	createRoomsTechTagTest(t, room)
}

func TestGetRoomsTechTags(t *testing.T) {
	n := 5
	room := createRoomsTest(t)

	for i := 0; i < n; i++ {
		createRoomsTechTagTest(t, room)
	}

	tags, err := testQueries.GetRoomsTechTags(context.Background(), room.RoomID)
	require.NoError(t, err)
	require.NotEmpty(t, tags)
	require.Len(t, tags, n)
}
