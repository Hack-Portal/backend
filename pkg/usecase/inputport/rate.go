package inputport

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
)

type RateUsecase interface {
	CreateRateEntry(ctx context.Context, body repository.CreateRateEntitiesParams) error
	ListRateEntry(ctx context.Context, id string, query domain.ListRequest) ([]repository.RateEntity, error)
	ListAccountRate(ctx context.Context, args domain.ListRequest) (result []domain.AccountRateResponse, err error)
}
