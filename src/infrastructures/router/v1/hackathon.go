package v1

import (
	"temp/src/adapter/controllers"
	"temp/src/adapter/gateways"
	"temp/src/adapter/presenters"
	"temp/src/usecase"
)

func (g *ginRouter) hackathonRouter() {
	hc := controllers.NewHackathonController(
		usecase.NewHackathon(
			presenters.NewHackathonPresenter(),
			gateways.NewHackathonGateway(g.db, g.app),
		),
		presenters.NewErrorPresenter(),
	)
	g.e.POST("/hackathon", hc.Create)
}
