package dai

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/src/datastructure/models"
)

type StatusTagDai interface {
	Create(ctx context.Context, statusTag *models.StatusTag) (id int64, err error)
	FindAll(ctx context.Context) (statusTags []*models.StatusTag, err error)
	FindById(ctx context.Context, id int64) (statusTag *models.StatusTag, err error)
	Update(ctx context.Context, statusTag *models.StatusTag) (id int64, err error)
	// TODO: Deleteする際にすでに割り当てられているStatusTagがある場合の一貫性をどうするかを検討する必要があるため保留
	// Delete(ctx context.Context, id int64) (err error)
}
