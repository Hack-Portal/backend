package interactors

import (
	"context"

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

func (hi *HackathonInteractor) CreateHackathonProposal(ctx context.Context, in request.CreateHackathonProposal) (int, *response.CreateHackathonProposal) {
	defer newrelic.FromContext(ctx).StartSegment("CreateHackathonProposal-usecase").End()

	// ハッカソン提案を作成
	id, err := hi.HackathonProposalOutput.Create(ctx, in.URL)
	if err != nil {
	}

	return hi.HackathonProposalOutput.PresentCreateHackathonProposal(ctx, &ports.OutputCreateHackathonProposalData{
		Error: nil,
		Response: &response.CreateHackathonProposal{
			HackathonProposalID: id,
			URL:                 in.URL,
		},
	})
}
