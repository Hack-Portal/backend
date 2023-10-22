package rotuer

import (
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
)

type Router interface {
}

type router struct {
	l   logger.Logger
	gin *gin.Engine
}

func NewServer(l logger.Logger, engine *gin.Engine) Router {
	serve := &router{
		gin: engine,
		l:   l,
	}

	serve.Account()

	return serve
}
