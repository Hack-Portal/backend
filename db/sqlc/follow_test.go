package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// フォローを追加する
func createFollowTest(t *testing.T, to, from Accounts) Follows {
	// アカウント追加のパラメタを満たす
	arg := CreateFollowParams{
		// 送り元ユーザID
		ToUserID: to.UserID,
		// 送り先ユーザID
		FromUserID: from.UserID,
	}
	// 実行する
	follow, err := testQueries.CreateFollow(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, follow)

	require.Equal(t, arg.ToUserID, follow.ToUserID)
	require.Equal(t, arg.FromUserID, follow.FromUserID)
	return follow
}

// Createのテスト
func TestCreateFollow(t *testing.T) {
	// ２つのアカウントを追加
	toAccount := createAccountTest(t)
	fromAccount := createAccountTest(t)

	createFollowTest(t, toAccount, fromAccount)
}

func TestRemoveFollow(t *testing.T) {
	n := 5
	toAccount := createAccountTest(t)
	var lastFollow Follows

	for i := 0; i < n; i++ {
		fromAccount := createAccountTest(t)
		lastFollow = createFollowTest(t, toAccount, fromAccount)
	}

	err := testQueries.RemoveFollow(context.Background(), RemoveFollowParams{
		ToUserID:   lastFollow.ToUserID,
		FromUserID: lastFollow.FromUserID,
	})

	require.NoError(t, err)

	listBookmarks, err := testQueries.ListFollowByToUserID(context.Background(), toAccount.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, listBookmarks)
	require.Len(t, listBookmarks, n-1)

	for _, follow := range listBookmarks {
		require.NotEmpty(t, follow)
		require.NotEqual(t, follow, lastFollow)
	}

}
