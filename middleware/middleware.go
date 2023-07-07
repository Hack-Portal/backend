package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/server"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
)

func AuthMiddleware(server *server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
		}
	}
}
