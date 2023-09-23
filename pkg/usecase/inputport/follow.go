package inputport

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/response"
)

type FollowUsecase interface {
	CreateFollow(ctx context.Context, body repository.CreateFollowsParams) (result response.FollowResponse, err error)
	RemoveFollow(ctx context.Context, body repository.DeleteFollowsParams) error
	GetFollowByID(ctx context.Context, ID string, mode bool) (result []response.FollowResponse, err error)
}
