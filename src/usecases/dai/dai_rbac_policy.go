package dai

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
)

type RBACPolicyDai interface {
	FindRoleByRole(ctx context.Context, role int) ([]*models.CasbinPolicy, error)
	FindRoleByPath(ctx context.Context, path string) ([]*models.CasbinPolicy, error)
	FindRoleByPathAndMethod(ctx context.Context, path, method string) ([]*models.CasbinPolicy, error)

	Create(ctx context.Context, policy []*models.RbacPolicy) ([]int64, error)
	FindAll(ctx context.Context) ([]*models.RbacPolicy, error)
	DeleteByID(ctx context.Context, id int64) error
	DeleteAll(ctx context.Context) error
}
