package dai

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
)

type HackathonStatusDai interface {
	Create(ctx context.Context, HackathonID string, hackathonStatus []int64) error
	FindAll(ctx context.Context, HackathonID []string) ([]*models.JoinedStatusTag, error)
	Delete(ctx context.Context, HackathonID string) error
}
