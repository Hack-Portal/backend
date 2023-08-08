package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	controller "github.com/hackhack-Geek-vol6/backend/pkg/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/gateways/repository/transaction"
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
