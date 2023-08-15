package route

import (
	"time"

	"github.com/gin-gonic/gin"
	controller "github.com/hackhack-Geek-vol6/backend/pkg/adapter/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
)

func NewRateRouter(env *bootstrap.Env, timeout time.Duration, store transaction.Store, group gin.IRoutes) {
	rateController := controller.RateController{
		RateUsecase: usecase.NewRateUsercase(store, timeout),
		Env:         env,
	}
	group.GET("/rate", rateController.ListAccountRate)
	group.GET("/rate/:account_id", rateController.ListRate)
	group.POST("/rate/:account_id", rateController.CreateRate)
}
