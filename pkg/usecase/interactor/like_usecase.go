package usecase

import (
	"context"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
)

type likeUsecase struct {
	store          transaction.Store
	contextTimeout time.Duration
}

func NewLikeUsercase(store transaction.Store, timeout time.Duration) inputport.LikeUsecase {
	return &likeUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (bu *likeUsecase) CreateLike(ctx context.Context, body repository.CreateLikesParams) (domain.BookmarkResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()
	bookmark, err := bu.store.CreateLikes(ctx, body)
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

func (bu *likeUsecase) GetLike(ctx context.Context, id string, query domain.ListRequest) (result []domain.BookmarkResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	bookmarks, err := bu.store.ListLikesByID(ctx, id)
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

func (bu *likeUsecase) RemoveLike(ctx context.Context, accountID string, opus int32) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	_, err := bu.store.DeleteLikesByID(ctx, repository.DeleteLikesByIDParams{AccountID: accountID, Opus: opus})
	return err
}
