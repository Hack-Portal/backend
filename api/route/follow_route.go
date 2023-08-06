package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/api/controller"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/usecase"
)

// フォローのルーティングを定義する
func NewFollowRouter(env *bootstrap.Env, timeout time.Duration, store db.Store, group *gin.RouterGroup) {
	followRepository := repository.NewFollowRepository(store, domain.CollectionFollow)
	followController := controller.FollowController{
		FollowUsecase: usecase.NewFollowUsercase(followRepository, timeout),
		Env:           env,
	}

	group.POST("/accounts/:id/follow", followController.GetFollow)
	group.POST("/accounts/:id/follow", followController.CreateFollow)
	group.DELETE("/acccounts/:id/follow	", followController.RemoveFollow)
}
