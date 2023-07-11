package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/util/token"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationClaimsKey  = "authorization_claim"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 1 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		accessToken := fields[0]
		hCS, err := token.JwtDecode.DecomposeFB(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		payload, err := token.JwtDecode.DecodeClaimFB(hCS[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set(AuthorizationClaimsKey, payload)
		ctx.Next()
	}
}
