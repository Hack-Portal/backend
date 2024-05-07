package interactors

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/hperror"
	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type HackathonProposalInteractor struct {
	HackathonProposal       dai.HackathonProposalDai
	HackathonProposalOutput ports.HackathonProposalOutputBoundary
}

func NewHackathonProposalInteractor(hackathonProposalDai dai.HackathonProposalDai, hackathonProposalOutput ports.HackathonProposalOutputBoundary) ports.HackathonProposalInputBoundary {
	return &HackathonProposalInteractor{
		HackathonProposal:       hackathonProposalDai,
		HackathonProposalOutput: hackathonProposalOutput,
	}
}

func (hi *HackathonProposalInteractor) CreateHackathonProposal(ctx context.Context, in request.CreateHackathonProposal) (int, *response.CreateHackathonProposal) {
	defer newrelic.FromContext(ctx).StartSegment("CreateHackathonProposal-usecase").End()

	if in.URL == "" {
		return hi.HackathonProposalOutput.PresentCreateHackathonProposal(ctx, &ports.OutputCreateHackathonProposalData{
			Error:    hperror.ErrFieldRequired,
			Response: nil,
		})
	}
	// ハッカソン提案を作成
	id, err := hi.HackathonProposal.Create(ctx, in.URL)
	if err != nil {
		return hi.HackathonProposalOutput.PresentCreateHackathonProposal(ctx, &ports.OutputCreateHackathonProposalData{
			Error:    err,
			Response: nil,
		})
	}

	return hi.HackathonProposalOutput.PresentCreateHackathonProposal(ctx, &ports.OutputCreateHackathonProposalData{
		Error: nil,
		Response: &response.CreateHackathonProposal{
			HackathonProposalID: int(id),
			URL:                 in.URL,
		},
	})
}
