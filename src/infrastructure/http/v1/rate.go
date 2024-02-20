package rotuer

import (
	"github.com/hackhack-Geek-vol6/backend/src/controllers"
)

func (r *router) rate() {
	rc := controllers.NewRateController(r.l, r.store)

	r.gin.GET("/rate", rc.ListAccountRate)
	r.gin.GET("/rate/:account_id", rc.ListRate)
	r.gin.POST("/rate/:account_id", rc.CreateRate)
}
