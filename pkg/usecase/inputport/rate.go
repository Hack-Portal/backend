package inputport

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
)

type RateUsecase interface {
	CreateRateEntry(ctx context.Context, body repository.CreateRateEntriesParams) error
	ListRateEntry(ctx context.Context, id string, query domain.ListRateParams) ([]repository.RateEntity, error)
}
