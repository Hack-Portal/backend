package db

import (
	"context"
	"log"
	"testing"

	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccountTx(t *testing.T) {
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
	log.Println(args.AccountTechTag)
	log.Println(args.AccountFrameworkTag)

	store := NewStore(testDB)

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
}