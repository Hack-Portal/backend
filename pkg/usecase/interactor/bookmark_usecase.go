package usecase

import (
	"context"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
)

type bookmarkUsecase struct {
	store          transaction.Store
	contextTimeout time.Duration
}

func NewBookmarkUsercase(store transaction.Store, timeout time.Duration) inputport.BookmarkUsecase {
	return &bookmarkUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (bu *bookmarkUsecase) CreateBookmark(ctx context.Context, body repository.CreateBookmarksParams) (domain.BookmarkResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()
	bookmark, err := bu.store.CreateBookmarks(ctx, body)
	if err != nil {
		return domain.BookmarkResponse{}, err
	}

	result, err := bu.store.GetHackathonByID(ctx, bookmark.Opus)
	if err != nil {
		return domain.BookmarkResponse{}, err
	}

	return domain.BookmarkResponse{
		HackathonID: result.HackathonID,
		Name:        result.Name,
		Icon:        result.Icon.String,
		Description: result.Description,
		Link:        result.Link,
		Expired:     result.Expired,
		StartDate:   result.StartDate,
		Term:        result.Term,
	}, nil
}

func (bu *bookmarkUsecase) GetBookmarks(ctx context.Context, id string, query domain.ListBookmarkRequestQueries) (result []domain.BookmarkResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	bookmarks, err := bu.store.ListBookmarksByID(ctx, id)
	if err != nil {
		return
	}
	for _, bookmark := range bookmarks {
		hackathon, err := bu.store.GetHackathonByID(ctx, bookmark.Opus)
		if err != nil {
			return nil, err
		}
		result = append(result, domain.BookmarkResponse{
			HackathonID: hackathon.HackathonID,
			Name:        hackathon.Name,
			Icon:        hackathon.Icon.String,
			Description: hackathon.Description,
			Link:        hackathon.Link,
			Expired:     hackathon.Expired,
			StartDate:   hackathon.StartDate,
			Term:        hackathon.Term,
		})
	}
	return
}

func (bu *bookmarkUsecase) RemoveBookmark(ctx context.Context, accountID string, opus int32) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	_, err := bu.store.DeleteBookmarksByID(ctx, repository.DeleteBookmarksByIDParams{AccountID: accountID, Opus: opus})
	return err
}
