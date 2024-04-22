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
			gateways.NewHackathonGateway(r.db),
			gateways.NewHackathonStatusGateway(r.db),
			gateways.NewCloudflareR2(
				config.Config.Buckets.Bucket,
				r.client,
				config.Config.Buckets.Expired,
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
