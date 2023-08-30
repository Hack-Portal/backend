package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/util/jwt"
)

const (
	AuthorizationHeaderKey     = "dbauthorization"
	AuthorizationType          = "dbauthorization_type"
	AuthorizationClaimsKey     = "authorization_claim"
	AuthorizationKeyNotInclude = "authorization_not_include"
)

func CheckJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.Set(AuthorizationKeyNotInclude, true)
		} else {
			DecodeJwt(ctx, authorizationHeader, false)
		}

	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
		log.Println("test", authorizationHeader)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		DecodeJwt(ctx, authorizationHeader, false)
	}
}

func DecodeJwt(ctx *gin.Context, header string, auth bool) {
	fields := strings.Fields(header)
	if len(fields) < 1 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
		return
	}

	accessToken := fields[0]
	payload, err := jwt.ValidJWTtoken(accessToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	ctx.Set(AuthorizationKeyNotInclude, auth)
	ctx.Set(AuthorizationClaimsKey, payload)
	ctx.Next()
}
