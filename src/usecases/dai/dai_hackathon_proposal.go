package dai

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/models"
)

type HackathonProposalDai interface {
	Create(ctx context.Context, url string) (*models.HackathonProposal, error)
	// Find(ctx context.Context, hackathonProposalID int) (*models.HackathonProposal, error)
	// FindAll(ctx context.Context, param request.ListHackathonProposal) ([]*models.HackathonProposal, error)
	// Update(ctx context.Context, hackathonProposal *models.HackathonProposal) error
	// Delete(ctx context.Context, hackathonProposalID int) error
}
