package route

import (
	"time"

	"github.com/gin-gonic/gin"
	controller "github.com/hackhack-Geek-vol6/backend/pkg/adapter/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
	tokens "github.com/hackhack-Geek-vol6/backend/pkg/util/token"
)

func NewUserRouter(env *bootstrap.Env, tokenMaker tokens.Maker, timeout time.Duration, store transaction.Store, group gin.IRoutes) {

	userController := controller.UserController{
		UserUsecase: usecase.NewUserUsercase(store, timeout, tokenMaker),
		Env:         env,
	}

	group.POST("/users", userController.CreateUser)
}
