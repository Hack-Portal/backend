package ports

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/datastructure/response"
)

type HackathonProposalInputBoundary interface {
	CreateHackathonProposal(ctx context.Context, in request.CreateHackathonProposal) (int, *response.CreateHackathonProposal)
	// GetHackathonProposal(ctx context.Context, hackathonProposalID int) (int, *response.GetHackathonProposal)
	// ListHackathonProposal(ctx context.Context, in request.ListHackathonProposal) (int, []*response.GetHackathonProposal)
	// UpdateHackathonProposal(ctx context.Context, in request.UpdateHackathonProposal) (int, *response.UpdateHackathonProposal)
	// DeleteHackathonProposal(ctx context.Context, hackathonProposalID int) (int, *response.DeleteHackathonProposal)
}

type HackathonProposalOutputBoundary interface {
	PresentCreateHackathonProposal(ctx context.Context, out *OutputCreateHackathonProposalData) (int, *response.CreateHackathonProposal)
	// PresentGetHackathonProposal(ctx context.Context, out *OutputGetHackathonProposalData) (int, *response.GetHackathonProposal)
	// PresentListHackathonProposal(ctx context.Context, out *OutputListHackathonProposalData) (int, []*response.GetHackathonProposal)
	// PresentUpdateHackathonProposal(ctx context.Context, out *OutputUpdateHackathonProposalData) (int, *response.UpdateHackathonProposal)
	// PresentDeleteHackathonProposal(ctx context.Context, out *OutputDeleteHackathonProposalData) (int, *response.DeleteHackathonProposal)
}

type OutputCreateHackathonProposalData struct {
	Error    error
	Response *response.CreateHackathonProposal
}

type OutputGetHackathonProposalData struct {
	Error    error
	Response *response.GetHackathonProposal
}

type OutputListHackathonProposalData struct {
	Error    error
	Response []*response.GetHackathonProposal
}

type OutputUpdateHackathonProposalData struct {
	Error    error
	Response *response.UpdateHackathonProposal
}

type OutputDeleteHackathonProposalData struct {
	Error    error
	Response *response.DeleteHackathonProposal
}
