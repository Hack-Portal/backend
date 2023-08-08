package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	controller "github.com/hackhack-Geek-vol6/backend/pkg/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/gateways/repository/transaction"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
)

// フォローのルーティングを定義する
func NewFollowRouter(env *bootstrap.Env, timeout time.Duration, store transaction.Store, group *gin.RouterGroup) {
	followController := controller.FollowController{
		FollowUsecase: usecase.NewFollowUsercase(store, timeout),
		Env:           env,
	}

	group.GET("/accounts/:user_id/follow", followController.GetFollow)
	group.POST("/accounts/:user_id/follow", followController.CreateFollow)
	group.DELETE("/acccounts/:user_id/follow	", followController.RemoveFollow)
}
