package dai

import (
	"context"
)

type HackathonProposalDai interface {
	Create(ctx context.Context, url string) (id int64, err error)
	// Find(ctx context.Context, hackathonProposalID int) (*models.HackathonProposal, error)
	// FindAll(ctx context.Context, param request.ListHackathonProposal) ([]*models.HackathonProposal, error)
	// Update(ctx context.Context, hackathonProposal *models.HackathonProposal) error
	// Delete(ctx context.Context, hackathonProposalID int) error
}
