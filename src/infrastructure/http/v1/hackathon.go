package rotuer

import (
	"gtihub.com/hackhack-Geek-vol6/backend/src/controllers"
)

func (r *router) hackathon() {
	hc := controllers.NewHackathonController(r.store, r.l)

	r.gin.POST("/hackathons", hc.CreateHackathon)
	r.gin.GET("/hackathons", hc.ListHackathons)
	r.gin.GET("/hackathons/:hackathon_id", hc.GetHackathonByID)
}
