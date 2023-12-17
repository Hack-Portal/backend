package echo

import (
	"github.com/Hack-Portal/backend/src/adapters/controllers"
	"github.com/Hack-Portal/backend/src/adapters/gateways"
	"github.com/Hack-Portal/backend/src/adapters/presenters"
	"github.com/Hack-Portal/backend/src/usecases/interactors"
)

func (es *echoServer) StatusTag() {
	sc := controllers.NewStatusTagController(
		interactors.NewStatusTagInteractor(
			gateways.NewStatusTagGateway(es.db),
			presenters.NewStatusTagPresenter(),
		),
	)

	// GetAllStatusTag
	es.v1.GET("/status_tags", sc.FindAllStatusTag)
	// get status tag by id
	es.v1.GET("/status_tags/:id", sc.FindByIdStatusTag)
	// create status tag
	es.v1.POST("/status_tags", sc.CreateStatusTag)
	// update status tag
	es.v1.PUT("/status_tags/:id", sc.UpdateStatusTag)
}
