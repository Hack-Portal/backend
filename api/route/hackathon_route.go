package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/api/controller"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/usecase"
)

func NewHackathonRouter(env *bootstrap.Env, timeout time.Duration, store db.Store, group *gin.RouterGroup) {

	hackathonRepository := repository.NewHackathonRepository(store, domain.CollectionHackathon)
	hackathonController := controller.HackathonController{
		HackathonUsecase: usecase.NewHackathonUsercase(hackathonRepository, timeout),
		Env:              env,
	}

	group.POST("/hackathons", hackathonController.CreateHackathon)
	group.GET("/hackathons", hackathonController.ListHackathons)
	group.GET("/hackathons/:hackathon_id", hackathonController.GetHackathon)
}
