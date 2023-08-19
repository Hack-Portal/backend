package transaction

import (
	"context"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	"github.com/stretchr/testify/require"
)

func randomRoom(t *testing.T, store *SQLStore) (domain.CreateRoomParam, repository.Account, repository.Room) {
	_, hackathon := randomHachathon(t, store)
	_, account := randomAccount(t, store)
	arg := domain.CreateRoomParam{
		RoomID:      util.RandomString(10),
		Title:       util.RandomString(10),
		Description: util.RandomString(10),
		HackathonID: hackathon.HackathonID,
		MemberLimit: int32(util.Random(5)),
		OwnerID:     account.AccountID,
		IncludeRate: false,
	}
	room, err := store.CreateRoomTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, room)
	return arg, account, room
}

func TestCreateRoomTx(t *testing.T) {
	store := NewStore(testDB, App)

	arg, owner, room := randomRoom(t, store)

	roomAccounts, err := store.GetRoomsAccountsByID(context.Background(), room.RoomID)
	require.NoError(t, err)
	require.NotEmpty(t, roomAccounts)
	require.Len(t, roomAccounts, 1)

	require.Equal(t, owner.AccountID, roomAccounts[0].AccountID.String)
	require.Equal(t, true, roomAccounts[0].IsOwner)

	require.Equal(t, arg.RoomID, room.RoomID)
	require.Equal(t, arg.Title, room.Title)
	require.Equal(t, arg.Description, room.Description)
	require.Equal(t, arg.HackathonID, room.HackathonID)
	require.Equal(t, arg.MemberLimit, room.MemberLimit)
	require.Equal(t, arg.IncludeRate, room.IncludeRate)
	require.Equal(t, false, room.IsDelete)
	require.NotZero(t, room.CreateAt)
	require.NotZero(t, room.UpdateAt)
}

func TestUpdateRoomTx(t *testing.T) {
	store := NewStore(testDB, App)
	_, hackathon := randomHachathon(t, store)
	_, owner, room := randomRoom(t, store)

	arg := domain.UpdateRoomParam{
		RoomID:      room.RoomID,
		Title:       util.RandomString(10),
		Description: util.RandomString(10),
		HackathonID: hackathon.HackathonID,
		MemberLimit: int32(util.Random(5)) + 1,
		OwnerEmail:  owner.Email,
	}

	newRoom, err := store.UpdateRoomTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, newRoom)

	require.Equal(t, arg.RoomID, newRoom.RoomID)
	require.Equal(t, arg.Title, newRoom.Title)
	require.Equal(t, arg.Description, newRoom.Description)
	require.Equal(t, arg.HackathonID, newRoom.HackathonID)
	require.Equal(t, arg.MemberLimit, newRoom.MemberLimit)
	require.Equal(t, false, newRoom.IsDelete)
	require.NotZero(t, newRoom.CreateAt)
	require.NotZero(t, newRoom.UpdateAt)
}

func TestDeleteRoomTx(t *testing.T) {
	store := NewStore(testDB, App)
	_, owner, room := randomRoom(t, store)

	arg := domain.DeleteRoomParam{
		OwnerEmail: owner.Email,
		RoomID:     room.RoomID,
	}

	require.NoError(t, store.DeleteRoomTx(context.Background(), arg))
}

func TestAddAccountInRoom(t *testing.T) {
	store := NewStore(testDB, App)
	_, owner, room := randomRoom(t, store)
	_, account := randomAccount(t, store)

	testCases := []struct {
		name        string
		arg         domain.AddAccountInRoomParam
		checkResult func(t *testing.T, err error)
	}{
		{
			name: "success",
			arg: domain.AddAccountInRoomParam{
				AccountID: account.AccountID,
				RoomID:    room.RoomID,
			},
			checkResult: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "fail duplicate account",
			arg: domain.AddAccountInRoomParam{
				AccountID: owner.AccountID,
				RoomID:    room.RoomID,
			},
			checkResult: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		tc.checkResult(t, store.AddAccountInRoom(context.Background(), tc.arg))
	}
}
