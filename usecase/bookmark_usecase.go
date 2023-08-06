package usecase

import (
	"context"
	"time"

	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/domain"
)

type bookmarkUsecase struct {
	store          db.Store
	contextTimeout time.Duration
}

func NewBookmarkUsercase(store db.Store, timeout time.Duration) domain.BookmarkUsecase {
	return &bookmarkUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (bu *bookmarkUsecase) CreateBookmark(ctx context.Context, body db.CreateBookmarkParams) (domain.BookmarkResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()
	bockmark, err := bu.store.CreateBookmark(ctx, body)
	if err != nil {
		return domain.BookmarkResponse{}, err
	}

	result, err := bu.store.GetHackathonByID(ctx, bockmark.HackathonID)
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

	bookmarks, err := bu.store.ListBookmarkByUserID(ctx, id)
	if err != nil {
		return
	}
	for _, bookmark := range bookmarks {
		hackathon, err := bu.store.GetHackathonByID(ctx, bookmark.HackathonID)
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

func (bu *bookmarkUsecase) RemoveBookmark(ctx context.Context, userID string, hackathonID int32) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	_, err := bu.store.SoftRemoveBookmark(ctx, db.SoftRemoveBookmarkParams{UserID: userID, HackathonID: hackathonID})
	return err
}