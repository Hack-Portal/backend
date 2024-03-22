package presenters

import (
	"context"
	"net/http"

	"github.com/Hack-Portal/backend/src/datastructure/hperror"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type HackathonProposalPresenter struct {
}

func NewHackathonProposalPresenter() ports.HackathonProposalOutputBoundary {
	return &HackathonProposalPresenter{}
}

func (s *HackathonProposalPresenter) PresentCreateHackathonProposal(ctx context.Context, out *ports.OutputCreateHackathonProposalData) (int, *response.CreateHackathonProposal) {
	defer newrelic.FromContext(ctx).StartSegment("PresentCreateHackathonProposal-presenter").End()

	if out.Error != nil {
		switch out.Error {
		case hperror.ErrFieldRequired:
			return http.StatusBadRequest, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	return http.StatusCreated, out.Response
}
