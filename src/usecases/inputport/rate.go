package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/src/domain/request"
	"github.com/hackhack-Geek-vol6/backend/src/domain/response"
	"github.com/hackhack-Geek-vol6/backend/src/repository"
)

type RateUsecase interface {
	CreateRateEntry(ctx context.Context, body repository.CreateRateEntitiesParams) error
	ListRateEntry(ctx context.Context, id string, query request.ListRequest) ([]repository.RateEntity, error)
	ListAccountRate(ctx context.Context, args request.ListRequest) (result []response.AccountRate, err error)
}
