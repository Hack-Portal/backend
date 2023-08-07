package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	controller "github.com/hackhack-Geek-vol6/backend/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/gateways/repository/transaction"
	usecase "github.com/hackhack-Geek-vol6/backend/usecase/interactor"
)

// フォローのルーティングを定義する
func NewFollowRouter(env *bootstrap.Env, timeout time.Duration, store transaction.Store, group *gin.RouterGroup) {
	followController := controller.FollowController{
		FollowUsecase: usecase.NewFollowUsercase(store, timeout),
		Env:           env,
	}

	group.GET("/accounts/:id/follow", followController.GetFollow)
	group.POST("/accounts/:id/follow", followController.CreateFollow)
	group.DELETE("/acccounts/:id/follow	", followController.RemoveFollow)
}
