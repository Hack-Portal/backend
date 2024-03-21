package ports

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/datastructure/response"
)

type HackathonProposalInputBoundary interface {
	CreateHackathonProposal(ctx context.Context, in request.CreateHackathonProposal) (int, *response.CreateHackathonProposal)
}

type HackathonProposalOutputBoundary interface {
	PresentCreateHackathonProposal(ctx context.Context, out *OutputCreateHackathonProposalData) (int, *response.CreateHackathonProposal)
}

type OutputCreateHackathonProposalData struct {
	Error    error
	Response *response.CreateHackathonProposal
}
