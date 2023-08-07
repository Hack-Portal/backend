package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/internal/controller/v1"
	"github.com/hackhack-Geek-vol6/backend/internal/usecase"
	"github.com/hackhack-Geek-vol6/backend/pkg/repository"
)

// アカウントのルーティングを定義する
func NewAccountRouter(env *bootstrap.Env, timeout time.Duration, store repository.Store, group *gin.RouterGroup) {
	accountController := controller.AccountController{
		AccountUsecase: usecase.NewAccountUsercase(store, timeout),
		Env:            env,
	}

	group.POST("/accounts", accountController.CreateAccount)
	group.GET("/accounts/:user_id", accountController.GetAccount)
	group.PUT("/accounts/:user_id", accountController.UpdateAccount)
	group.DELETE("/acccounts/:user_id", accountController.DeleteAccount)
}
