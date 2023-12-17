package ports

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/src/datastructure/models"
	"github.com/hackhack-Geek-vol6/backend/src/datastructure/request"
	"github.com/hackhack-Geek-vol6/backend/src/datastructure/response"
)

type StatusTagInputBoundary interface {
	CreateStatusTag(ctx context.Context, in *request.CreateStatusTag) (int, *response.StatusTag)
	FindAllStatusTag(ctx context.Context) (int, []*response.StatusTag)
	FindByIdStatusTag(ctx context.Context, in *request.GetStatusTagByID) (int, *response.StatusTag)
	UpdateStatusTag(ctx context.Context, in *request.UpdateStatusTag) (int, *response.StatusTag)
	// TODO: Deleteする際にすでに割り当てられているStatusTagがある場合の一貫性をどうするかを検討する必要があるため保留
	// DeleteStatusTag(input *DeleteStatusTagInputData) (*DeleteStatusTagOutputData, error)
}

type StatusTagOutputBoundary interface {
	PresentCreateStatusTag(ctx context.Context, out *OutputCraeteStatusTagData) (int, *response.StatusTag)
	PresentFindAllStatusTag(ctx context.Context, out *OutputFindAllStatusTagData) (int, []*response.StatusTag)
	PresentFindByIdStatusTag(ctx context.Context, out *OutputFindByIdStatusTagData) (int, *response.StatusTag)
	PresentUpdateStatusTag(ctx context.Context, out *OutputUpdateStatusTagData) (int, *response.StatusTag)

	// TODO: Deleteする際にすでに割り当てられているStatusTagがある場合の一貫性をどうするかを検討する必要があるため保留
	// PresentDeleteStatusTag(ctx context.Context, out *outputDeleteStatusTagData) (int, *response.StatusTagResponse)
}

type OutputCraeteStatusTagData struct {
	Error    error
	Response *models.StatusTag
}

type OutputFindAllStatusTagData struct {
	Error    error
	Response []*models.StatusTag
}

type OutputFindByIdStatusTagData struct {
	Error    error
	Response *models.StatusTag
}
type OutputUpdateStatusTagData struct {
	Error    error
	Response *models.StatusTag
}
