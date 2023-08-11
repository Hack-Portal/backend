package transaction_test

import (
	"context"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	tx "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/test/repository"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	"github.com/stretchr/testify/require"
)

func TestCreateAccountTx(t *testing.T) {
	store := NewStore(testDB, &fb.App{})

	techTags := ListTechTagTest(t)
	frameworks := test.listFrameworkTest(t)
	techTagIds := util.RandomSelection(len(techTags), 10)
	frameworkIds := util.RandomSelection(len(frameworks), 10)
	args := tx.CreateAccountTxParams{
		Accounts: repository.Account{
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
	var accountTechTags []repository.TechTag
	var accountFrameworks []repository.Framework

	for _, techtagid := range techTagIds {
		accountTechTag, err := store.GetTechTagByID(context.Background(), techtagid)
		require.NoError(t, err)
		require.NotEmpty(t, accountTechTag)
		accountTechTags = append(accountTechTags, accountTechTag)
	}

	for _, frameworkId := range frameworkIds {
		accountFramework, err := store.GetFrameworksByID(context.Background(), frameworkId)
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
}
