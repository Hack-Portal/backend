package rotuer

import (
	"gtihub.com/hackhack-Geek-vol6/backend/src/controllers"
)

func (r *router) etc() {
	ec := controllers.NewEtcController(r.store, r.l)

	r.gin.GET("/health", ec.Health)
	r.gin.GET("/frameworks", ec.ListFrameworks)
	r.gin.GET("/locates", ec.ListLocation)
	r.gin.GET("/tech_tags", ec.ListTechTag)
	r.gin.GET("/status_tags", ec.ListStatusTag)
	r.gin.GET("/roles", ec.ListRoles)
}
