package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/internal/controller/v1"
	"github.com/hackhack-Geek-vol6/backend/internal/usecase"
	"github.com/hackhack-Geek-vol6/backend/pkg/repository"
)

func NewHackathonRouter(env *bootstrap.Env, timeout time.Duration, store repository.Store, group *gin.RouterGroup) {
	hackathonController := controller.HackathonController{
		HackathonUsecase: usecase.NewHackathonUsercase(store, timeout),
		Env:              env,
	}

	group.POST("/hackathons", hackathonController.CreateHackathon)
	group.GET("/hackathons", hackathonController.ListHackathons)
	group.GET("/hackathons/:hackathon_id", hackathonController.GetHackathon)
}
