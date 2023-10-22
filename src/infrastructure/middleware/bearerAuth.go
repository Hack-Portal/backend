package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/util/jwt"
)

const (
	AuthorizationHeaderKey     = "dbauthorization"
	AuthorizationClaimsKey     = "authorization_claim"
	AuthorizationKeyNotInclude = "authorization_not_include"
)

func (m *middleware) JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.Set(AuthorizationKeyNotInclude, true)
		} else {
			m.decodeJwt(ctx, authorizationHeader, false)
		}
	}
}

func (m *middleware) decodeJwt(ctx *gin.Context, header string, auth bool) {
	fields := strings.Fields(header)
	if len(fields) < 1 {
		m.l.Debug("invalid authorization header format")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
		return
	}

	accessToken := fields[0]
	payload, err := jwt.ValidJWTtoken(accessToken)
	if err != nil {
		m.l.Debugf("invalid jwt token: %v", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
	}
	ctx.Set(AuthorizationKeyNotInclude, auth)
	ctx.Set(AuthorizationClaimsKey, payload)
	ctx.Next()
}
