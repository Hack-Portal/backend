package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/internal/controller/v1"
	"github.com/hackhack-Geek-vol6/backend/internal/usecase"
	"github.com/hackhack-Geek-vol6/backend/pkg/repository"
)

func NewEtcRouter(env *bootstrap.Env, timeout time.Duration, store repository.Store, group *gin.RouterGroup) {
	etcController := controller.EtcController{
		EtcUsecase: usecase.NewEtcUsercase(store, timeout),
		Env:        env,
	}

	group.GET("/ping", etcController.Ping)
	group.GET("/locates", etcController.ListLocation)
	group.GET("/tech_tags", etcController.ListTechTags)
	group.GET("/frameworks", etcController.ListFrameworks)
}
