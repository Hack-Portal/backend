package dai

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
)

type HackathonDai interface {
	Create(ctx context.Context, hackathon *models.Hackathon) error
}
