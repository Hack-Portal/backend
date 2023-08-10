package repository_test

import (
	"context"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/stretchr/testify/require"
)

func createBookmarkTest(t *testing.T, account repository.Account) repository.Bookmark {
	hackathon := createHackathonTest(t)
	arg := repository.CreateBookmarksParams{
		HackathonID: hackathon.HackathonID,
		AccountID:   account.AccountID,
	}
	bookmark, err := testQueries.CreateBookmarks(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bookmark)

	require.Equal(t, arg.HackathonID, bookmark.HackathonID)
	require.Equal(t, arg.AccountID, bookmark.AccountID)
	return bookmark
}

func TestCreateBookmark(t *testing.T) {
	account := CreateAccountTest(t)
	createBookmarkTest(t, account)
}

func TestListBookmark(t *testing.T) {
	n := 5
	account := CreateAccountTest(t)

	var bookmarks []repository.Bookmark
	for i := 0; i < n; i++ {
		bookmark := createBookmarkTest(t, account)
		bookmarks = append(bookmarks, bookmark)
	}
	results, err := testQueries.ListBookmarksByID(context.Background(), account.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, results)
	require.Len(t, results, n)

	for index, result := range results {
		require.NotEmpty(t, result)
		require.Equal(t, bookmarks[index], result)
	}

}

func TestSoftRemoveBookmark(t *testing.T) {
	n := 5
	account := CreateAccountTest(t)
	var lastBookMark repository.Bookmark
	for i := 0; i < n; i++ {
		lastBookMark = createBookmarkTest(t, account)
	}

	_, err := testQueries.DeleteBookmarksByID(context.Background(), repository.DeleteBookmarksByIDParams{
		AccountID:   lastBookMark.AccountID,
		HackathonID: lastBookMark.HackathonID,
	})
	require.NoError(t, err)

	listBookmark, err := testQueries.ListBookmarksByID(context.Background(), account.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, listBookmark)

	require.Len(t, listBookmark, n-1)

	for _, bookmark := range listBookmark {
		require.NotEqual(t, bookmark, lastBookMark)
	}
}
