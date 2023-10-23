package rotuer

import (
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	"github.com/hackhack-Geek-vol6/backend/src/repository"
)

type Router interface {
}

type router struct {
	l     logger.Logger
	gin   *gin.Engine
	store repository.SQLStore
}

func NewServer(l logger.Logger, engine *gin.Engine) Router {
	serve := &router{
		gin: engine,
		l:   l,
	}

	serve.account()
	serve.etc()
	serve.follow()

	return serve
}
