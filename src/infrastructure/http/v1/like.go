package rotuer

import (
	"github.com/hackhack-Geek-vol6/backend/src/controllers"
)

func (r *router) like() {
	lc := controllers.NewLikeController(r.l, r.store)

	r.gin.POST("/like", lc.CreateLikeEntry)
	r.gin.GET("/like/:user_id", lc.GetLikeEntry)
	r.gin.DELETE("/like/:user_id", lc.DeleteLikeEntry)
}
