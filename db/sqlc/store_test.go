package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccountTx(t *testing.T) {
	store := NewStore(testDB)

	techTags := listTechTagTest(t)
	frameworks := listFrameworkTest(t)
	techTagIds := util.RandomSelection(len(techTags), 10)
	frameworkIds := util.RandomSelection(len(frameworks), 10)
	args := CreateAccountTxParams{
		Accounts: Accounts{
			UserID:     util.RandomString(8),
			Username:   util.RandomString(8),
			LocateID:   int32(util.Random(47)),
			Rate:       0,
			Email:      util.RandomEmail(),
			ShowRate:   true,
			ShowLocate: true,
		},
		AccountTechTag:      techTagIds,
		AccountFrameworkTag: frameworkIds,
	}
	var accountTechTags []TechTags
	var accountFrameworks []Frameworks

	for _, techtagid := range techTagIds {
		accountTechTag, err := store.GetTechTag(context.Background(), techtagid)
		require.NoError(t, err)
		require.NotEmpty(t, accountTechTag)
		accountTechTags = append(accountTechTags, accountTechTag)
	}

	for _, frameworkId := range frameworkIds {
		accountFramework, err := store.GetFrameworks(context.Background(), frameworkId)
		require.NoError(t, err)
		require.NotEmpty(t, accountFramework)
		accountFrameworks = append(accountFrameworks, accountFramework)
	}

	result, err := store.CreateAccountTx(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, args.Accounts.UserID, result.Account.UserID)
	require.Equal(t, args.Accounts.Username, result.Account.Username)
	require.Equal(t, args.Accounts.Email, result.Account.Email)
	require.Equal(t, args.Accounts.Rate, result.Account.Rate)
	require.Equal(t, args.Accounts.LocateID, result.Account.LocateID)
	require.Equal(t, args.Accounts.ShowLocate, result.Account.ShowLocate)
	require.Equal(t, args.Accounts.ShowRate, result.Account.ShowRate)

	require.Len(t, result.AccountTechTags, len(techTagIds))
	require.Len(t, result.AccountFrameworks, len(frameworkIds))

	require.Equal(t, result.AccountTechTags, accountTechTags)
	require.Equal(t, result.AccountFrameworks, accountFrameworks)
}

func TestCreateRoomTx(t *testing.T) {
	store := NewStore(testDB)

	techTags := listTechTagTest(t)
	frameworks := listFrameworkTest(t)
	techTagIds := util.RandomSelection(len(techTags), 10)
	frameworkIds := util.RandomSelection(len(frameworks), 10)

	user := createAccountTest(t)
	roomID, err := uuid.NewRandom()
	require.NoError(t, err)
	require.NotEmpty(t, roomID)

	args := CreateRoomTxParams{
		Rooms: Rooms{
			RoomID:      roomID,
			HackathonID: createHackathonTest(t).HackathonID,
			Title:       util.RandomString(8),
			Description: util.RandomString(100),
			MemberLimit: 5,
		},
		UserID:          user.UserID,
		RoomsTechTags:   techTagIds,
		RoomsFrameworks: frameworkIds,
	}
	var roomsTechTags []TechTags
	var roomsFrameworks []Frameworks

	for _, techtagid := range techTagIds {
		accountTechTag, err := store.GetTechTag(context.Background(), techtagid)
		require.NoError(t, err)
		require.NotEmpty(t, accountTechTag)
		roomsTechTags = append(roomsTechTags, accountTechTag)
	}

	for _, frameworkId := range frameworkIds {
		accountFramework, err := store.GetFrameworks(context.Background(), frameworkId)
		require.NoError(t, err)
		require.NotEmpty(t, accountFramework)
		roomsFrameworks = append(roomsFrameworks, accountFramework)
	}

	result, err := store.CreateRoomTx(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.NotZero(t, result.RoomID)
	require.Equal(t, args.HackathonID, result.HackathonID)
	require.Equal(t, args.Title, result.Title)
	require.Equal(t, args.Description, result.Description)
	require.Equal(t, args.MemberLimit, result.MemberLimit)

	require.Equal(t, args.UserID, result.RoomsAccounts[0].UserID.String)
	require.Equal(t, user.Username, result.RoomsAccounts[0].Username.String)
	require.Equal(t, user.Icon, result.RoomsAccounts[0].Icon)

	require.Len(t, result.RoomsTechTags, len(techTagIds))
	require.Len(t, result.RoomsFrameworks, len(frameworkIds))

	require.Equal(t, result.RoomsTechTags, roomsTechTags)
	require.Equal(t, result.RoomsFrameworks, roomsFrameworks)
}

// func TestCreateHackathonTx(t *testing.T) {
// 	store := NewStore(testDB)
// 	statusTags, err := store.ListStatusTags(context.Background())

// 	statusTagIds := util.RandomSelection(len(statusTags), 2)

// 	args := CreateHackathonTxParams{
// 		Hackathons: Hackathons{
// 			Name:        util.RandomString(8),
// 			Icon:        []byte(util.RandomString(8)),
// 			Description: util.RandomString(100),
// 			Link:        util.RandomString(16),
// 			Expired:     time.Now().Add(time.Hour * 100),
// 			StartDate:   time.Now().Add(time.Hour * 200),
// 			Term:        int32(util.Random(100)),
// 		},
// 		HackathonStatusTag: statusTagIds,
// 	}
// 	var hackathonsStatusTags []StatusTags

// 	result, err := store.CreateHackathonTx(context.Background(), args)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, result)

// 	require.NotZero(t, result.HackathonID)
// 	require.Equal(t, args.Name, result.Name)
// 	require.Equal(t, args.Icon, result.Icon)
// 	require.Equal(t, args.Description, result.Description)
// 	require.Equal(t, args.Link, result.Link)
// 	require.NotZero(t, result.Expired)
// 	require.NotZero(t, result.StartDate)
// 	require.Equal(t, args.Term, result.Term)

// 	require.Len(t, result.HackathonStatusTags, len(statusTagIds))
// 	require.Equal(t, result.HackathonStatusTags, hackathonsStatusTags)
// }
