package inputport

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
)

type PastworkUsecase interface {
	CreatePastWork(ctx context.Context, arg domain.CreatePastWorkParams, image []byte) (result domain.PastWorkResponse, err error)
	GetPastWork(ctx context.Context, opus int32) (result domain.PastWorkResponse, err error)
	ListPastWork(ctx context.Context, query domain.ListRequest) (result []domain.ListPastWorkResponse, err error)
	UpdatePastWork(ctx context.Context, body repository.UpdatePastWorksByIDParams) (result domain.PastWorkResponse, err error)
	DeletePastWork(ctx context.Context, args repository.DeletePastWorksByIDParams) error
}
