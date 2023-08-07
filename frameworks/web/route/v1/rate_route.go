package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/internal/controller/v1"
	"github.com/hackhack-Geek-vol6/backend/internal/usecase"
	"github.com/hackhack-Geek-vol6/backend/pkg/repository"
)

func NewRateRouter(env *bootstrap.Env, timeout time.Duration, store repository.Store, group *gin.RouterGroup) {
	rateController := controller.RateController{
		RateUsecase: usecase.NewRateUsercase(store, timeout),
		Env:         env,
	}
	group.GET("/accounts/:user_id/rate", rateController.ListRate)
	group.POST("/accounts/:user_id/rate", rateController.CreateRate)
}
