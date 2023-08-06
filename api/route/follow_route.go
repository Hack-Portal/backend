package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
)

// フォローのルーティングを定義する
func NewFollowRouter(env *bootstrap.Env, timeout time.Duration, store db.Store, group *gin.RouterGroup) {
	followRepository := repository.NewFollowRepository(store, domain.CollectionFollow)
	followController := controller.FollowController{
		FollowUsecase: usecase.NewFollowUsercase(followRepository, timeout),
		Env:           env,
	}

	group.POST("/accounts", followController.CreateAccount)

	group.GET("/accounts/:id", followController.GetAccount)
	group.PUT("/accounts/:id", followController.UpdateAccount)
	group.DELETE("/acccounts/:id", followController.DeleteAccount)
}
