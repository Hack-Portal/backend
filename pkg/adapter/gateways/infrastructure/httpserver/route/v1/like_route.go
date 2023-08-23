package route

import (
	"time"

	"github.com/gin-gonic/gin"
	controller "github.com/hackhack-Geek-vol6/backend/pkg/adapter/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
)

// ブックマークのルーティングを定義する
func NewLikeRouter(env *bootstrap.Env, timeout time.Duration, store transaction.Store, group gin.IRoutes) {
	LikeController := controller.LikeController{
		LikeUsecase: usecase.NewLikeUsercase(store, timeout),
		Env:         env,
	}
	group.POST("/pastworks/:opus/like", LikeController.CreateLike)
	group.GET("/pastworks/:opus/like/:user_id", LikeController.ListLike)
	group.DELETE("/pastworks/:opus/like/:user_id", LikeController.RemoveLike)
}
