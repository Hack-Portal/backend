package dai

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
)

// HackathonStatusDai はHackathonStatusに関するデータアクセスインターフェース
type HackathonStatusDai interface {
	FindAll(ctx context.Context, HackathonID []string) ([]*models.JoinedStatusTag, error)
	Delete(ctx context.Context, HackathonID string) error
}
