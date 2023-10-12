package transaction

import (
	"context"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/params"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/etc"
	"github.com/stretchr/testify/require"
)

func randomRoom(t *testing.T) (params.CreateRoom, repository.Account, repository.Room) {
	_, hackathon := randomHachathon(t)
	_, account := randomAccount(t)
	arg := params.CreateRoom{
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
	arg, owner, room := randomRoom(t)

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
	_, hackathon := randomHachathon(t)
	_, owner, room := randomRoom(t)
	_, account := randomAccount(t)

	testCases := []struct {
		name        string
		arg         params.UpdateRoom
		checkResult func(t *testing.T, arg params.UpdateRoom, room repository.Room, err error)
	}{
		{
			name: "success",
			arg: params.UpdateRoom{
				RoomID:      room.RoomID,
				Title:       util.RandomString(10),
				Description: util.RandomString(10),
				HackathonID: hackathon.HackathonID,
				MemberLimit: int32(util.Random(5)) + 1,
				OwnerEmail:  owner.Email,
			},
			checkResult: func(t *testing.T, arg params.UpdateRoom, room repository.Room, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, room)

				require.Equal(t, arg.RoomID, room.RoomID)
				require.Equal(t, arg.Title, room.Title)
				require.Equal(t, arg.Description, room.Description)
				require.Equal(t, arg.HackathonID, room.HackathonID)
				require.Equal(t, arg.MemberLimit, room.MemberLimit)
				require.Equal(t, false, room.IsDelete)
				require.NotZero(t, room.CreateAt)
				require.NotZero(t, room.UpdateAt)
			},
		}, {
			name: "fail not owner",
			arg: params.UpdateRoom{
				RoomID:      room.RoomID,
				Title:       util.RandomString(10),
				Description: util.RandomString(10),
				HackathonID: hackathon.HackathonID,
				MemberLimit: int32(util.Random(5)) + 1,
				OwnerEmail:  account.Email,
			},
			checkResult: func(t *testing.T, arg params.UpdateRoom, room repository.Room, err error) {
				require.Error(t, err)
			},
		}, {
			name: "success - close room",
			arg: params.UpdateRoom{
				RoomID:     room.RoomID,
				IsClosing:  true,
				OwnerEmail: owner.Email,
			},
			checkResult: func(t *testing.T, arg params.UpdateRoom, room repository.Room, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, room)

				require.Equal(t, arg.RoomID, room.RoomID)
				require.Equal(t, false, room.IsDelete)
				require.NotZero(t, room.CreateAt)
				require.NotZero(t, room.UpdateAt)
			},
		},
	}
	for _, tc := range testCases {
		newRoom, err := store.UpdateRoomTx(context.Background(), tc.arg)
		tc.checkResult(t, tc.arg, newRoom, err)
	}
}

func TestDeleteRoomTx(t *testing.T) {
	_, owner, room := randomRoom(t)

	arg := params.DeleteRoom{
		OwnerEmail: owner.Email,
		RoomID:     room.RoomID,
	}

	require.NoError(t, store.DeleteRoomTx(context.Background(), arg))
}

func TestAddAccountInRoom(t *testing.T) {
	_, owner, room := randomRoom(t)
	_, account := randomAccount(t)

	testCases := []struct {
		name        string
		arg         params.AddAccountInRoom
		checkResult func(t *testing.T, err error)
	}{
		{
			name: "success",
			arg: params.AddAccountInRoom{
				AccountID: account.AccountID,
				RoomID:    room.RoomID,
			},
			checkResult: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "fail duplicate account",
			arg: params.AddAccountInRoom{
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

func TestCloseRoom(t *testing.T) {
	// アカウントを6つくらい生成する
	var accountID []string
	_, owner, room := randomRoom(t)
	for i := 0; i < 6; i++ {
		_, account := randomAccount(t)
		accountID = append(accountID, account.AccountID)
	}

	arg := params.CloseRoom{
		RoomID:    room.RoomID,
		AccountID: account.AccountID,
	}

	require.NoError(t, store.CloseRoom(context.Background(), arg))
}

func TestAddRoomAccountRoleByID(t *testing.T) {
	_, owner, room := randomRoom(t)
	_, account := randomAccount(t)
	// Roleを追加する
	arg := params.RoomAccountRole{
		RoomID:    room.RoomID,
		AccountID: account.AccountID,
	}
	testCases := []struct {
		name        string
		arg         params.RoomAccountRole
		checkResult func(t *testing.T, err error)
	}{
		{
			name: "success",
			arg: params.RoomAccountRole{
				AccountID: account.AccountID,
				RoomID:    room.RoomID,
				RoleID:    []int32{1, 2},
			},
			checkResult: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "fail duplicate account",
			arg: params.RoomAccountRole{
				AccountID: owner.AccountID,
				RoomID:    room.RoomID,
				RoleID:    []int32{1, 2},
			},
			checkResult: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		tc.checkResult(t, store.AddRoomAccountRoleByID(context.Background(), tc.arg))
	}
	require.NoError(t, store.AddRoomAccountRoleByID(context.Background(), arg))
}

func TestUpdateRoomsAccountRoleByID(t *testing.T) {
	_, owner, room := randomRoom(t)
	_, account := randomAccount(t)

	arg := params.UpdateRoomsAccountRoleByID{
		AccountID: account.AccountID,
		RoomID:    room.RoomID,
		IsOwner:   true,
	}

	require.NoError(t, store.UpdateRoomsAccountRoleByID(context.Background(), arg))
}
