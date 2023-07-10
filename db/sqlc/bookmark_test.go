package db

import (
	"context"
	"testing"
)

func createBookmarkTest(t *testing.T,account Accounts) {
	hackathon := createHackathonTest(t)
	
	bookmark , err := testQueries.CreateBookmark(context.Background()),
}
