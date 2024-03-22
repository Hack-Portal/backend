package v1

import (
	"github.com/Hack-Portal/backend/src/adapters/controllers"
	"github.com/Hack-Portal/backend/src/adapters/gateways"
	"github.com/Hack-Portal/backend/src/adapters/presenters"
	"github.com/Hack-Portal/backend/src/usecases/interactors"
)

func (r *v1router) hackathonProposal() {
	hackathonProposal := r.v1.Group("/hackathons-proposal")

	// DI
	hc := controllers.NewHackathonProposalController(
		interactors.NewHackathonProposalInteractor(
			gateways.NewHackathonProposalGateway(r.db, r.cache),
			presenters.NewHackathonProposalPresenter(),
		),
	)

	hackathonProposal.POST("", hc.CreateHackathonProposal)
	// hackathonProposal.GET("", hc.ListHackathonProposals)
	// hackathonProposal.PUT("/:hackathon_proposal_id", hc.UpdateHackathonProposal)
	// hackathonProposal.DELETE("/:hackathon_proposal_id", hc.DeleteHackathonProposal)

}
