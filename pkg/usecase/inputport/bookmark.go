package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/gateways/repository/datasource"
)

type BookmarkUsecase interface {
	CreateBookmark(ctx context.Context, body repository.CreateBookmarkParams) (domain.BookmarkResponse, error)
	GetBookmarks(ctx context.Context, id string, query domain.ListBookmarkRequestQueries) (result []domain.BookmarkResponse, err error)
	RemoveBookmark(ctx context.Context, userID string, hackathonID int32) error
}
