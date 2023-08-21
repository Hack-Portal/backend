package route

import (
	"time"

	"github.com/gin-gonic/gin"
	controller "github.com/hackhack-Geek-vol6/backend/pkg/adapter/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
)

// アカウントのルーティングを定義する
func NewAccountRouter(env *bootstrap.Env, timeout time.Duration, store transaction.Store, group gin.IRoutes) {
	accountController := controller.AccountController{
		AccountUsecase: usecase.NewAccountUsercase(store, timeout),
		Env:            env,
	}
	group.POST("/accounts", accountController.CreateAccount)
	group.GET("/accounts/:account_id", accountController.GetAccount)
	group.PUT("/accounts/:account_id", accountController.UpdateAccount)
	group.DELETE("/accounts/:account_id", accountController.DeleteAccount)
}
