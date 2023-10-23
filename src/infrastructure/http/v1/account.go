package rotuer

import (
	"gtihub.com/hackhack-Geek-vol6/backend/src/controllers"
)

func (r *router) account() {
	ac := controllers.NewAccountController(r.store, r.l)

	r.gin.POST("/account", ac.CreateAccount)

	r.gin.GET("/account/:id", ac.GetAccountByID)
	r.gin.PUT("/account/:id", ac.UpdateAccount)
	r.gin.DELETE("/account/:id", ac.DeleteAccount)

	r.gin.GET("/account/:id/join-room", ac.GetJoinRoom)
}
