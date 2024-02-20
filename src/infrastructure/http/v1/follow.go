package rotuer

import (
	"gtihub.com/hackhack-Geek-vol6/backend/src/controllers"
)

func (r *router) follow() {
	fc := controllers.NewFollowController(r.store, r.l)

	r.gin.GET("/accounts/:account_id/follow", fc.GetFollow)
	r.gin.POST("/accounts/:account_id/follow", fc.CreateFollow)
	r.gin.DELETE("/accounts/:account_id/follow", fc.DeleteFollow)
}
