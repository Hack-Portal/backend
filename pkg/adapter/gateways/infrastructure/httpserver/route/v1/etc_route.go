package route

import (
	"time"

	"github.com/gin-gonic/gin"
	controller "github.com/hackhack-Geek-vol6/backend/pkg/adapter/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewEtcRouter(env *bootstrap.Env, timeout time.Duration, store transaction.Store, group *gin.RouterGroup) {
	etcController := controller.EtcController{
		EtcUsecase: usecase.NewEtcUsercase(store, timeout),
		Env:        env,
	}

	group.GET("/ping", etcController.Ping)
	group.GET("/locates", etcController.ListLocation)
	group.GET("/tech_tags", etcController.ListTechTags)
	group.GET("/frameworks", etcController.ListFrameworks)
}

func setupSwagger(group *gin.RouterGroup) {
	group.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
