package repository_test

import (
	"context"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	"github.com/stretchr/testify/require"
)

func createAccountTagsTest(t *testing.T, account repository.Account) repository.AccountTag {
	tags := listTechTagTest(t)
	techs := util.Random(len(tags) - 1)

	arg := repository.CreateAccountTagsParams{
		UserID:    account.UserID,
		TechTagID: tags[techs].TechTagID,
	}
	accountTags, err := testQueries.CreateAccountTags(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, accountTags)

	require.Equal(t, arg.UserID, accountTags.UserID)
	require.Equal(t, arg.TechTagID, accountTags.TechTagID)

	return accountTags
}

func TestCreateAccountTang(t *testing.T) {
	account := createAccountTest(t)
	createAccountTagsTest(t, account)
}

func TestGetAccountTag(t *testing.T) {
	n := 5
	account := createAccountTest(t)
	for i := 0; i < n; i++ {
		createAccountTagsTest(t, account)
	}

	tags, err := testQueries.ListAccountTagsByUserID(context.Background(), account.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, tags)
	require.Len(t, tags, n)
}

func TestDeleteAccountTagByUserID(t *testing.T) {
	n := 5
	account := createAccountTest(t)
	for i := 0; i < n; i++ {
		createAccountTagsTest(t, account)
	}

	err := testQueries.DeleteAccountTagsByUserID(context.Background(), account.UserID)
	require.NoError(t, err)
	accountFramework, err := testQueries.ListAccountTagsByUserID(context.Background(), account.UserID)
	require.NoError(t, err)
	require.Len(t, accountFramework, 0)
	for _, af := range accountFramework {
		require.Empty(t, af)
	}
}
