package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/api/controller"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/usecase"
)

func NewRateRouter(env *bootstrap.Env, timeout time.Duration, store db.Store, group *gin.RouterGroup) {

	rateRepository := repository.NewRateRepository(store, domain.CollectionRate)
	rateController := controller.RateController{
		RateUsecase: usecase.NewRateUsercase(rateRepository, timeout),
		Env:         env,
	}
	group.GET("/accounts/:user_id/rate", rateController.ListRate)
	group.POST("/accounts/:user_id/rate", rateController.CreateRate)
}
