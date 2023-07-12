package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createFollowTest(t *testing.T, account Accounts) Follows {
	arg := CreateFollowParams{
		ToUserID:   account.UserID,
		FromUserID: account.UserID,
	}
	follow, err := testQueries.CreateFollow(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, follow)

	require.Equal(t, arg.ToUserID, follow.ToUserID)
	require.Equal(t, arg.FromUserID, follow.FromUserID)
	return follow
}

func TestCreateFollow(t *testing.T) {
	account := createAccountTest(t)
	createFollowTest(t, account)
}

func TestRemoveFollow(t *testing.T) {
	n := 5
	account := createAccountTest(t)
	var lastFollow Follows
	for i := 0; i < n; i++ {
		createFollowTest(t, account)
	}

	err := testQueries.RemoveFollow(context.Background(), RemoveFollowParams{
		ToUserID:   lastFollow.ToUserID,
		FromUserID: lastFollow.FromUserID,
	})
	require.NoError(t, err)

	listBookmark, err := testQueries.ListFollowByToUserID(context.Background(), account.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, listBookmark)
	require.Len(t, listBookmark, n)

	for _, follow := range listBookmark {
		require.NotEmpty(t, follow)
		require.NotEqual(t, follow, lastFollow)
	}

}
