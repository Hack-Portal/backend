package dai

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
)

// RoleDai はRoleに関するデータアクセスインターフェース
type RoleDai interface {
	Create(ctx context.Context, roleStore *models.Role) (id int64, err error)
	FindAll(ctx context.Context) (roleStores []*models.Role, err error)
	FindByID(ctx context.Context, id int64) (roleStore *models.Role, err error)
	Update(ctx context.Context, roleStore *models.Role) (id int64, err error)
}
