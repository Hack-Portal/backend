package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/api/controller"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/usecase"
)

func NewEtcRouter(env *bootstrap.Env, timeout time.Duration, store db.Store, group *gin.RouterGroup) {

	etcRepository := repository.NewEtcRepository(store, domain.CollectionEtc)
	etcController := controller.EtcController{
		EtcUsecase: usecase.NewEtcUsercase(etcRepository, timeout),
		Env:        env,
	}

	group.GET("/ping", etcController.Ping)
	group.GET("/locates", etcController.ListLocation)
	group.GET("/tech_tags", etcController.ListTechTags)
	group.GET("/frameworks", etcController.ListFrameworks)
}
