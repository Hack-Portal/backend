package repository_test

import (
	"context"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/stretchr/testify/require"
)

func createFollowTest(t *testing.T, to, from repository.Account) repository.Follow {
	// アカウント追加のパラメタを満たす
	arg := repository.CreateFollowsParams{
		// 送り元ユーザID
		ToAccountID: to.AccountID,
		// 送り先ユーザID
		FromAccountID: from.AccountID,
	}
	// 実行する
	follow, err := testQueries.CreateFollows(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, follow)

	require.Equal(t, arg.ToAccountID, follow.ToAccountID)
	require.Equal(t, arg.FromAccountID, follow.FromAccountID)
	return follow
}

// Createのテスト
func TestCreateFollow(t *testing.T) {
	// ２つのアカウントを追加
	toAccount := CreateAccountTest(t)
	fromAccount := CreateAccountTest(t)

	createFollowTest(t, toAccount, fromAccount)
}

func TestRemoveFollow(t *testing.T) {
	n := 5
	toAccount := CreateAccountTest(t)
	var lastFollow repository.Follow

	for i := 0; i < n; i++ {
		fromAccount := CreateAccountTest(t)
		lastFollow = createFollowTest(t, toAccount, fromAccount)
	}

	err := testQueries.DeleteFollows(context.Background(), repository.DeleteFollowsParams{
		ToAccountID:   lastFollow.ToAccountID,
		FromAccountID: lastFollow.FromAccountID,
	})

	require.NoError(t, err)

	listBookmarks, err := testQueries.ListFollowsByToUserID(context.Background(), toAccount.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, listBookmarks)
	require.Len(t, listBookmarks, n-1)

	for _, follow := range listBookmarks {
		require.NotEmpty(t, follow)
		require.NotEqual(t, follow, lastFollow)
	}

}
