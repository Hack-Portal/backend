package inputport

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/params"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/request"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/response"
)

type PastworkUsecase interface {
	CreatePastWork(ctx context.Context, arg params.CreatePastWork, image []byte) (result response.PastWork, err error)
	GetPastWork(ctx context.Context, opus int32) (result response.PastWork, err error)
	ListPastWork(ctx context.Context, query request.ListRequest) (result []response.ListPastWork, err error)
	UpdatePastWork(ctx context.Context, body params.UpdatePastWork) (result response.PastWork, err error)
	DeletePastWork(ctx context.Context, args repository.DeletePastWorksByIDParams) error
}
