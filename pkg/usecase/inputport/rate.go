package inputport

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/request"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/response"
)

type RateUsecase interface {
	CreateRateEntry(ctx context.Context, body repository.CreateRateEntitiesParams) error
	ListRateEntry(ctx context.Context, id string, query request.ListRequest) ([]repository.RateEntity, error)
	ListAccountRate(ctx context.Context, args request.ListRequest) (result []response.AccountRateResponse, err error)
}
