package ports

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/datastructure/response"
)

// StatusTagInputBoundary はStatusTagのInputBoundary
type StatusTagInputBoundary interface {
	CreateStatusTag(ctx context.Context, in *request.CreateStatusTag) (int, *response.StatusTag)
	FindAllStatusTag(ctx context.Context) (int, []*response.StatusTag)
	FindByIDStatusTag(ctx context.Context, in *request.GetStatusTagByID) (int, *response.StatusTag)
	UpdateStatusTag(ctx context.Context, in *request.UpdateStatusTag) (int, *response.StatusTag)
	// TODO: Deleteする際にすでに割り当てられているStatusTagがある場合の一貫性をどうするかを検討する必要があるため保留
	// DeleteStatusTag(input *DeleteStatusTagInputData) (*DeleteStatusTagOutputData, error)
}

// StatusTagOutputBoundary はStatusTagのOutputBoundary
type StatusTagOutputBoundary interface {
	PresentCreateStatusTag(ctx context.Context, out OutputBoundary[*models.StatusTag]) (int, *response.StatusTag)
	PresentFindAllStatusTag(ctx context.Context, out OutputBoundary[[]*models.StatusTag]) (int, []*response.StatusTag)
	PresentFindByIDStatusTag(ctx context.Context, out OutputBoundary[*models.StatusTag]) (int, *response.StatusTag)
	PresentUpdateStatusTag(ctx context.Context, out OutputBoundary[*models.StatusTag]) (int, *response.StatusTag)

	// TODO: Deleteする際にすでに割り当てられているStatusTagがある場合の一貫性をどうするかを検討する必要があるため保留
	// PresentDeleteStatusTag(ctx context.Context, out *outputDeleteStatusTagData) (int, *response.StatusTagResponse)
}
