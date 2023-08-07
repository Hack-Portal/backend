package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/domain"
	repository "github.com/hackhack-Geek-vol6/backend/gateways/repository/datasource"
)

type FollowUsecase interface {
	CreateFollow(ctx context.Context, body repository.CreateFollowParams) (result domain.FollowResponse, err error)
	RemoveFollow(ctx context.Context, body repository.RemoveFollowParams) error
	GetFollowByToID(ctx context.Context, ID string) (result []domain.FollowResponse, err error)
}
