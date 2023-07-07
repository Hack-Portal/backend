package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
