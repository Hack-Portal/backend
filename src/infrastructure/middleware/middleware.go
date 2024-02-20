package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
)

type Middleware interface {
	Logger() gin.HandlerFunc
	Apm() gin.HandlerFunc
	JwtAuth() gin.HandlerFunc
}

type middleware struct {
	l logger.Logger
}

func New(l logger.Logger) Middleware {
	return &middleware{l}
}
