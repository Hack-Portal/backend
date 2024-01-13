package dai

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
)

// UsersDai はUserに関するデータアクセスインターフェース
type UsersDai interface {
	Create(ctx context.Context, user *models.User) (id string, err error)
	FindAll(ctx context.Context) (users []*models.User, err error)
	FindByID(ctx context.Context, id string) (user *models.User, err error)
	Update(ctx context.Context, user *models.User) (id string, err error)
	Delete(ctx context.Context, id string) (err error)
}
