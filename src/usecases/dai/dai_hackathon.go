package dai

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
)

type HackathonDai interface {
	Create(ctx context.Context, hackathon *models.Hackathon, hackathonStatus []int64) error
	Find(ctx context.Context, hackathonID string) (*models.Hackathon, error)
	FindAll(ctx context.Context, param FindAllParams) ([]*models.Hackathon, error)
	Delete(ctx context.Context, hackathonID string) error
}

type FindAllParams struct {
	Limit  int
	Offset int

	// タグ
	Tags []int64
	// 新着かどうか？
	New bool
	// 期間が長いかどうか？
	LongTerm bool
	// 締め切りが近いかどうか？
	NearDeadline bool
}
