package v1

import (
	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/Hack-Portal/backend/src/adapters/controllers"
	"github.com/Hack-Portal/backend/src/adapters/gateways"
	"github.com/Hack-Portal/backend/src/adapters/presenters"
	"github.com/Hack-Portal/backend/src/usecases/interactors"
	"github.com/labstack/echo/v4"
)

func (r *v1router) hackathon() {
	hackathon := r.v1.Group("/hackathons")

	// DI
	hc := controllers.NewHackathonController(
		interactors.NewHackathonInteractor(
			gateways.NewHackathonGateway(r.db, r.cache),
			gateways.NewHackathonStatusGateway(r.db, r.cache),
			gateways.NewCloudflareR2(
				config.Config.Buckets.Bucket,
				r.client,
				r.cache,
				gateways.WithPresignLinkExpired(config.Config.Buckets.Expired),
			),
			interactors.NewDiscordNotifyInteractor(
				gateways.NewDiscordChannelGateway(r.db),
				gateways.NewDiscordServerRegistryGateways(r.db),
				gateways.NewDiscordNotifyGateway(r.session),
			),
			presenters.NewHackathonPresenter(),
		),
	)

	hackathon.POST("", hc.CreateHackathon)
	hackathon.GET("", hc.ListHackathons)
	hackathon.PUT("/:hackathon_id", func(c echo.Context) error {
		return c.String(500, "not implemented")
	})
	hackathon.DELETE("/:hackathon_id", hc.DeleteHackathon)
}
