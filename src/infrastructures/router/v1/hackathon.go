package v1

import (
	"github.com/hackhack-Geek-vol6/backend/src/adapter/controllers"
	"github.com/hackhack-Geek-vol6/backend/src/adapter/gateways"
	"github.com/hackhack-Geek-vol6/backend/src/adapter/presenters"
	"github.com/hackhack-Geek-vol6/backend/src/usecase"
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
