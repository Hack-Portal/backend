package route

import (
	"time"

	"github.com/gin-gonic/gin"
	controller "github.com/hackhack-Geek-vol6/backend/pkg/adapter/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
)

func NewRateRouter(env *bootstrap.Env, timeout time.Duration, store transaction.Store, group *gin.RouterGroup) {
	rateController := controller.RateController{
		RateUsecase: usecase.NewRateUsercase(store, timeout),
		Env:         env,
	}
	group.GET("/accounts/:user_id/rate", rateController.ListRate)
	group.POST("/accounts/:user_id/rate", rateController.CreateRate)
}
