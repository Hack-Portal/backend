package inputport

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
)

type LikeUsecase interface {
	CreateLike(ctx context.Context, body repository.CreateLikesParams) (domain.LikeResponse, error)
	GetLike(ctx context.Context, id string, query domain.ListRequest) (result []domain.LikeResponse, err error)
	RemoveLike(ctx context.Context, userID string, hackathonID int32) error
}
