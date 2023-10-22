package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/src/domain/request"
	"github.com/hackhack-Geek-vol6/backend/src/repository"
)

type LikeUsecase interface {
	CreateLike(ctx context.Context, body repository.CreateLikesParams) (repository.Like, error)
	GetLike(ctx context.Context, id string, query request.ListRequest) (result []repository.Like, err error)
	RemoveLike(ctx context.Context, userID string, hackathonID int32) error
}
