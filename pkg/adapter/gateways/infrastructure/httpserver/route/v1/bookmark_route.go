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
func NewBookmarkRouter(env *bootstrap.Env, timeout time.Duration, store transaction.Store, group *gin.RouterGroup) {
	bookmarkController := controller.BookmarkController{
		BookmarkUsecase: usecase.NewBookmarkUsercase(store, timeout),
		Env:             env,
	}
	group.POST("/bookmarks", bookmarkController.CreateBookmark)
	group.GET("/bookmarks/:user_id", bookmarkController.ListBookmark)
	group.DELETE("/bookmarks/:user_id", bookmarkController.RemoveBookmark)
}
