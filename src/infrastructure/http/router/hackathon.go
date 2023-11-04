package router

import (
	"github.com/hackhack-Geek-vol6/backend/src/adapters/controllers"
	"github.com/hackhack-Geek-vol6/backend/src/adapters/gateways"
	"github.com/hackhack-Geek-vol6/backend/src/adapters/presenters"
)

func (r *ginRouter) hackathonRouter() {
	hr := r.e.Group("/v1/hackathons").Use()

	hc := controllers.NewHackathonController(
		presenters.NewHackathonOutputBoundary(),
		gateways.NewHackathonGateway(r.db),
		gateways.NewFirebaseRepository(r.app),
	)

	hr.POST("/", hc.Create)
	hr.GET("/", hc.ReadAll)
	hr.DELETE("/", hc.DeleteAll)

	hr.PUT("/:id", hc.Update)
	hr.DELETE("/:id", hc.Delete)
}
