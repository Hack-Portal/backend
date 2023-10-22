package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/src/domain/response"
	"github.com/hackhack-Geek-vol6/backend/src/repository"
)

type FollowUsecase interface {
	CreateFollow(ctx context.Context, body repository.CreateFollowsParams) (result response.Follow, err error)
	RemoveFollow(ctx context.Context, body repository.DeleteFollowsParams) error
	GetFollowByID(ctx context.Context, ID string, mode bool) (result []response.Follow, err error)
}
