package route

import (
	"time"

	"github.com/gin-gonic/gin"
	controller "github.com/hackhack-Geek-vol6/backend/pkg/adapter/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
)

func NewPastWorkRouter(env *bootstrap.Env, timeout time.Duration, store transaction.Store, group gin.IRoutes) {
	PastWorkController := controller.PastWorkController{
		PastWorkUsecase: usecase.NewPastWorkUsercase(store, timeout),
		Env:             env,
	}
	group.POST("/pastworks", PastWorkController.CreatePastWork)
	group.GET("/pastworks", PastWorkController.ListPastWork)
	group.GET("/pastworks/:opus", PastWorkController.GetPastWork)
	group.PUT("/pastworks/:opus", PastWorkController.UpdatePastWork)
	group.DELETE("/pastworks/:opus", PastWorkController.DeletePastWork)
}
