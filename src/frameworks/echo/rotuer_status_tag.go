package echo

func (es *echoServer) StatusTag() {
	// GetAllStatusTag
	es.v1.GET("/status_tags", nil)
	// get status tag by id
	es.v1.GET("/status_tags/:id", nil)
	// create status tag
	es.v1.POST("/status_tags", nil)
	// update status tag
	es.v1.PUT("/status_tags/:id", nil)
}
