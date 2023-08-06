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
	group.GET("/accounts/:user_id", accountController.GetAccount)
	group.PUT("/accounts/:user_id", accountController.UpdateAccount)
	group.DELETE("/acccounts/:user_id", accountController.DeleteAccount)
}
