package inputport

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
)

type FollowUsecase interface {
	CreateFollow(ctx context.Context, body repository.CreateFollowParams) (result domain.FollowResponse, err error)
	RemoveFollow(ctx context.Context, body repository.RemoveFollowParams) error
	GetFollowByToID(ctx context.Context, ID string) (result []domain.FollowResponse, err error)
}
