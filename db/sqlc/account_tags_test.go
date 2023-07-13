package db

import (
	"context"
	"testing"

	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/stretchr/testify/require"
)

func createAccountTagsTest(t *testing.T, account Accounts) AccountTags {
	tags := listTechTagTest(t)
	techs := util.Random(len(tags) - 1)

	arg := CreateAccountTagsParams{
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
