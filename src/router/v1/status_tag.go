package v1

import (
	"github.com/Hack-Portal/backend/src/adapters/controllers"
	"github.com/Hack-Portal/backend/src/adapters/gateways"
	"github.com/Hack-Portal/backend/src/adapters/presenters"
	"github.com/Hack-Portal/backend/src/usecases/interactors"
)

func (r *v1router) statusTag() {
	statusTag := r.v1.Group("/status_tags")

	// DI
	sc := controllers.NewStatusTagController(
		interactors.NewStatusTagInteractor(
			gateways.NewStatusTagGateway(r.db, r.cache),
			presenters.NewStatusTagPresenter(),
		),
	)

	statusTag.POST("", sc.CreateStatusTag)
	statusTag.GET("", sc.FindAllStatusTag)
	statusTag.PUT("/:hackathon_id", sc.UpdateStatusTag)
}
