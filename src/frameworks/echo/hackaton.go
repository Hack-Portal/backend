package echo

import (
	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/Hack-Portal/backend/src/adapters/controllers"
	"github.com/Hack-Portal/backend/src/adapters/gateways"
	"github.com/Hack-Portal/backend/src/adapters/presenters"
	"github.com/Hack-Portal/backend/src/usecases/interactors"
)

func (es *echoServer) Hackathon() {
	hc := controllers.NewHackathonController(
		interactors.NewHackathonInteractor(
			gateways.NewHackathonGateway(es.db, es.redis),
			gateways.NewHackathonStatusGateway(es.db, es.redis),
			gateways.NewCloudflareR2(
				config.Config.Buckets.Bucket,
				es.client,
				es.redis,
				config.Config.Buckets.Expired,
			),
			presenters.NewHackathonPresenter(),
		),
	)

	// GetAllStatusTag
	es.v1.POST("/hackathons", hc.CreateHackathon)
	es.v1.GET("/hackathons", hc.ListHackathons)
	es.v1.GET("/hackathons/:hackathon_id", hc.GetHackathon)
	es.v1.DELETE("/hackathons/:hackathon_id", hc.DeleteHackathon)
}
