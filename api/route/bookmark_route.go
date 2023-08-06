package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/usecase"
)

// アカウントのルーティングを定義する
func NewBookmarkRouter(env *bootstrap.Env, timeout time.Duration, store db.Store, group *gin.RouterGroup) {
	bookmarkRepository := repository.NewBookmarkRepository(store, domain.CollectionAccount)
	bookmarkController := controller.BookmarkController{
		AccountUsecase: usecase.NewBookmarkUsercase(bookmarkRepository, timeout),
		Env:            env,
	}

	group.POST("/accounts", bookmarkController.CreateAccount)

	group.GET("/accounts/:id", bookmarkController.GetAccount)
	group.PUT("/accounts/:id", bookmarkController.UpdateAccount)
	group.DELETE("/acccounts/:id", bookmarkController.DeleteAccount)
}
