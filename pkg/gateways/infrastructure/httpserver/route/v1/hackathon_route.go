package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	controller "github.com/hackhack-Geek-vol6/backend/pkg/controllers/v1"
	"github.com/hackhack-Geek-vol6/backend/pkg/gateways/repository/transaction"
	usecase "github.com/hackhack-Geek-vol6/backend/pkg/usecase/interactor"
)

func NewHackathonRouter(env *bootstrap.Env, timeout time.Duration, store transaction.Store, group *gin.RouterGroup) {
	hackathonController := controller.HackathonController{
		HackathonUsecase: usecase.NewHackathonUsercase(store, timeout),
		Env:              env,
	}

	group.POST("/hackathons", hackathonController.CreateHackathon)
	group.GET("/hackathons", hackathonController.ListHackathons)
	group.GET("/hackathons/:hackathon_id", hackathonController.GetHackathon)
}
