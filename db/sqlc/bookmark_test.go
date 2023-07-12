package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createBookmarkTest(t *testing.T, account Accounts) Bookmarks {
	hackathon := createHackathonTest(t)
	arg := CreateBookmarkParams{
		HackathonID: hackathon.HackathonID,
		UserID:      account.UserID,
	}
	bookmark, err := testQueries.CreateBookmark(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bookmark)

	require.Equal(t, arg.HackathonID, bookmark.HackathonID)
	require.Equal(t, arg.UserID, bookmark.UserID)
	return bookmark
}

func TestCreateBookmark(t *testing.T) {
	account := createAccountTest(t)
	createBookmarkTest(t, account)
}

func TestListBookmark(t *testing.T) {
	n := 5
	account := createAccountTest(t)

	var bookmarks []Bookmarks
	for i := 0; i < n; i++ {
		bookmark := createBookmarkTest(t, account)
		bookmarks = append(bookmarks, bookmark)
	}
	results, err := testQueries.ListBookmarkByUserID(context.Background(), account.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, results)
	require.Len(t, results, n)

	for index, result := range results {
		require.NotEmpty(t, result)
		require.Equal(t, bookmarks[index], result)
	}

}

func TestRemoveBookmark(t *testing.T) {
	n := 5
	account := createAccountTest(t)
	var lastBookMark Bookmarks
	for i := 0; i < n; i++ {
		lastBookMark = createBookmarkTest(t, account)
	}

	err := testQueries.RemoveBookmark(context.Background(), RemoveBookmarkParams{
		UserID:      lastBookMark.UserID,
		HackathonID: lastBookMark.HackathonID,
	})
	require.NoError(t, err)

	listBookmark, err := testQueries.ListBookmarkByUserID(context.Background(), account.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, listBookmark)

	require.Len(t, listBookmark, n-1)

	for _, bookmark := range listBookmark {
		require.NotEqual(t, bookmark, lastBookMark)
	}
}
