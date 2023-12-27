package dai

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
)

type HackathonDai interface {
	Create(ctx context.Context, hackathon *models.Hackathon, hackathonStatus []int64) error
	Find(ctx context.Context, hackathonID string) (*models.Hackathon, error)
	FindAll(ctx context.Context, size, id int) ([]*models.Hackathon, error)
}
