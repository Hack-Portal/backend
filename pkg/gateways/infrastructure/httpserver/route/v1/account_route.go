package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	controller "github.com/hackhack-Geek-vol6/backend/pkg/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/gateways/repository/transaction"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
)

// アカウントのルーティングを定義する
func NewAccountRouter(env *bootstrap.Env, timeout time.Duration, store transaction.Store, group *gin.RouterGroup) {
	accountController := controller.AccountController{
		AccountUsecase: usecase.NewAccountUsercase(store, timeout),
		Env:            env,
	}

	group.POST("/accounts", accountController.CreateAccount)
	group.GET("/accounts/:user_id", accountController.GetAccount)
	group.PUT("/accounts/:user_id", accountController.UpdateAccount)
	group.DELETE("/acccounts/:user_id", accountController.DeleteAccount)
}
