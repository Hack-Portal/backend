package rotuer

import (
	"github.com/hackhack-Geek-vol6/backend/src/controllers"
)

func (r *router) pastWork() {
	pc := controllers.NewPastworkController(r.l, r.store)

	r.gin.POST("/pastwork", pc.CreatePastwork)
	r.gin.GET("/pastwork/", pc.ListPastwork)
	r.gin.GET("/pastwork/:opus", pc.GetPastwork)
	r.gin.PUT("/pastwork/:opus", pc.UpdatePastwork)
	r.gin.DELETE("/pastwork/:opus", pc.DeletePastwork)
}
