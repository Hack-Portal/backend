package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/api/controller"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/usecase"
)

// アカウントのルーティングを定義する
func NewAccountRouter(env *bootstrap.Env, timeout time.Duration, store db.Store, group *gin.RouterGroup) {
	accountRepository := repository.NewAccountRepository(store, domain.CollectionAccount)
	accountController := controller.AccountController{
		AccountUsecase: usecase.NewAccountUsercase(accountRepository, timeout),
		Env:            env,
	}

	group.POST("/accounts", accountController.CreateAccount)

	group.GET("/accounts/:id", accountController.GetAccount)
	group.PUT("/accounts/:id", accountController.UpdateAccount)
	group.DELETE("/acccounts/:id", accountController.DeleteAccount)

	// auth.GET("/acccounts/:id/follow")
	group.POST("/acccounts/:id/follow", accountController.CreateFollow)
	group.DELETE("/acccounts/:id/follow", accountController.RemoveFollow)
	// レート
	group.POST("/accounts/:id/rate", accountController.CreateRate)
	group.GET("/accounts/:id/rate", accountController.ListRate)
}
